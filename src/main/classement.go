package main
	
	import (
	"strconv"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"bufio"
	"strings"
	"regexp"
	"time"
	)
	
	type Classement struct {
	id int
	nom string 
	prenom string
	sexe string
	equipe string
	resultat int 
	}
	
	