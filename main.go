package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hapoon/gendao/log"
	"github.com/hapoon/gendao/repository"
	"github.com/hapoon/gendao/generator"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"path/filepath"
	"io/ioutil"
	"strings"
)

var (
	// DBHost is host name of database
	DBHost string
	// DBPort is port number of database
	DBPort int
	// DBUser is user name to access database
	DBUser string
	// DBPassword is password to access database
	DBPassword string
	// DBName is database namespace to access
	DBName string
	// DB is pointer of information for current DB connection
	DB *gorm.DB
	// VerboseMode is debug flag
	VerboseMode bool
	// Output is output directory
	Output string
	// Exclude is table names that exclude generating dao code.
	Exclude string
)

func main() {
	log.Init()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	prepareFlag()
	flag.Parse()
	if err := validateFlag(); err != nil {
		abend(err)
	}
	if VerboseMode {
		log.DebugMode = true
	}
	log.Debugf("host: %v port: %v user: %v password: %v dbname: %v\n", DBHost, DBPort, DBUser, DBPassword, DBName)
	args := flag.Args()
	log.Debugf("args: %v\n", args)

	prepareDB(DBHost, DBPort, DBUser, DBPassword, DBName)
	defer DB.Close()

	tables, err := showTables()
	if err != nil {
		abend(err)
	}

	columns := repository.Columns{}
	columns.New(DB)
	for _, table := range tables {
		if err := columns.FindAll(table); err != nil {
			abend(err)
		}
	}
	packageName := filepath.Base(Output)
	if Output == "" {
		Output, err = os.Getwd()
		if err != nil {
			abend(err)
		}
		packageName = "output"
		Output = filepath.Join(Output,packageName)
	}
	excludes := strings.Split(Exclude, ",")
	log.Infof("exclude tables %v\n", excludes)
	var daoList []string
	// create dao files
	for _, v := range columns.Data {
		if searchStrings(excludes, v.Name) {
			continue
		}
		filename := fmt.Sprintf("%s.go", v.Name)
		outputName := filepath.Join(Output,filename)
		daoName := generator.FormatDAOName(generator.ToCamelCaseFromSnakeCase(v.Name))
		daoList = append(daoList, daoName)
		// If output file exists, skip creating dao file.
		if _,err := os.Stat(outputName);err == nil {
			continue
		}
		g := generator.Generator{}
		g.GenerateHead(packageName)

		g.GenerateDAO(v)

		src := g.Format()
		err := ioutil.WriteFile(outputName, src, 0644)
		if err != nil {
			abend(err)
		}
		log.Infof("%s generated\n", outputName)
	}
	log.Debugf("daoList: %v\n", daoList)

	// create dao/misc.go
	createMisc(packageName, daoList)
}

func abend(err error) {
	log.Fatalln(err)
	os.Exit(2)
}

func prepareFlag() {
	// Database host
	flag.StringVar(&DBHost, "host", "", "Database host")
	flag.StringVar(&DBHost, "h", "", "Database host")
	// Database port
	flag.IntVar(&DBPort, "port", 0, "Database port number")
	flag.IntVar(&DBPort, "p", 0, "Database port number")
	// Database User
	flag.StringVar(&DBUser, "user", "", "Database user name")
	flag.StringVar(&DBUser, "u", "", "Database user name")
	// Database Password
	flag.StringVar(&DBPassword, "password", "", "Database password")
	flag.StringVar(&DBPassword, "pass", "", "Database password")
	// Database Name
	flag.StringVar(&DBName, "name", "", "Database name")
	flag.StringVar(&DBName, "n", "", "Database name")
	// Verbose
	flag.BoolVar(&VerboseMode, "verbose", false, "Verbose mode")
	// output
	flag.StringVar(&Output, "output", "", "Output directory; default current directory")
	flag.StringVar(&Output,"o", "", "Output directory; default current directory")
	// exclude
	flag.StringVar(&Exclude, "exclude", "", "exclude table names(separated by comma)")
	flag.StringVar(&Exclude, "e", "", "exclude table names(separated by comma)")
}

func prepareDB(host string, port int, user string, password string, dbname string) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		abend(err)
	}
	DB = db
}

func validateFlag() error {
	// validate database host
	if DBHost == "" {
		return errors.New("database host is required")
	}
	// validate database port
	if DBPort == 0 {
		return errors.New("database port is required")
	}
	// validate database user
	if DBUser == "" {
		return errors.New("database user is required")
	}
	// validate database password
	if DBPassword == "" {
		return errors.New("database password is required")
	}
	// validate database name
	if DBName == "" {
		return errors.New("database name is required")
	}
	return nil
}

func showTables() (result []string, err error) {
	rows, err := DB.Raw("SHOW TABLES").Rows()
	if err != nil {
		return
	}
	columns, err := rows.Columns()
	if err != nil {
		return
	}

	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			log.Debugln(columns[i], ": ", value)
			result = append(result, value)
		}
	}
	return
}

func searchStrings(str []string, search string) bool {
	for _,v := range str {
		if v == search {
			return true
		}
	}
	return false
}

func createMisc(packageName string, daoList []string) {
	outputName := filepath.Join(Output,"misc.go")
	g := generator.Generator{}
	g.GenerateHead(packageName)
	g.GenerateMISC(daoList)
	src := g.Format()
	err := ioutil.WriteFile(outputName, src, 0644)
	if err != nil {
		abend(err)
	}
	log.Infof("%s generated\n",outputName)	
}
