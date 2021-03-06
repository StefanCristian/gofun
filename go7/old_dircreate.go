package main

import (
        "fmt"
        //"strings"
	//"io/ioutil"
	"path/filepath"
	"path"
        "os"
	"runtime"
        //"errors"
	//. "github.com/kr/fs"
)

func init() {
	if len(os.Args) >= 2 {
		fmt.Printf("Error: You cannot add additional options or path names\n")
		os.Exit(1)
	}
}

func main() {
	fmt.Printf("Hello there, name a directory that you want to see if it exists or not:\n")
	var existenta string
	fmt.Scanf("%s",&existenta)
	//sz, err := fmt.Println(existenta)
	if path.IsAbs(existenta) == true {
	   //fmt.Printf("Yes, it is")
	   //continue
	   } else {
	   fmt.Printf("Error: This is not an absolute path\n")
	   os.Exit(1)
	   }

	zz, err := fmt.Println(path.Clean(existenta))
	if err != nil { 
	  fmt.Println(err)
	  } else {
	  fmt.Println(zz)
	  }

	if Exists(existenta) {
	   fmt.Printf("yay. it's there. something. Let's see if it's chroot-able, shall we?\n")
	   dirname := existenta + string(filepath.Separator)
	   if IfDirectory(dirname) == true {
	     d, err := os.Open(dirname)
	     if err != nil {
	   	//fmt.Println(err)
		os.Exit(1)
		}
		f, err := d.Readdir(-1)
		if err != nil {
		//fmt.Println(err)
		os.Exit(1)
		}
			//chroot prereqs
		   	for _, f := range f {
			//fmt.Println(f.Name(), f.Size())
			// I should add a mode  isregular ? with an `if`
				if f.Name() == "root" {
				fmt.Printf("oh hello root\n")
				   rootz := dirname + "root" + string(filepath.Separator)
				   edd, err := os.Open(rootz)
				     if err != nil {
				     fmt.Println(err)
				     os.Exit(1)
				     }
				     passf, err := edd.Readdir(-1)
				     if err != nil {
			 	     fmt.Println(err)
				     os.Exit(1)
				     }
				       for _, passf := range passf {
						fmt.Println(f.Name() + "/" + passf.Name())
				       }
				continue
				}
				if f.Name() == "bin" {
				fmt.Printf("oh hello bin\n")
				}
				// will add more verifies, and more professional also
			}
			fmt.Println(runtime.NumCgoCall()) //out of curiosity, heh
			//blah, err := exec.Command(equo, "login").Run()
	   }
	   
	   //if dirscan, err := ioutil.ReadDir(existenta); err != nil {
		//fmt.Println(err)
		//} else {
		//fmt.Println(dirscan)
		//}
		//IfDirectory(existenta)
	} else {
		var i string
		fmt.Printf("meh. it`s not there, Joe. Let`s start the creation procedure. Do you want to create it? Yes/No\n")
		fmt.Scanf("%s",&i)
		//i_yes := []string{"Y","yes","Yes","Ye","YES","YeS","Ye","y"}
		//i_no := []string{"No","N","NO","nO"}
		if i == "Yes" || i == "YES" || i == "Y" || i == "y" || i == "Ye" || i == "YE" || i == "yes" || i == "ye" {
                	fmt.Printf("Then name a directory name\n")
		        var dir string
			fmt.Scanf("%s",&dir)
			fmt.Println(os.Mkdir(dir, 22))
		} else {
			fmt.Printf("Have fun, then\n")
		}
	}
	IfDirectory(existenta)
	//ww := os.Mkdir(existenta, 22)

        //if fileinfo.IsDir(); err != nil {
        	//fmt.Printf("It's already created. Let's see if it's a file or a directory\n")
        //} else {
        	//fmt.Println(err)
        	//fmt.Printf("It's a already-made directory\n")
        //}
}


type Cmd struct {
	equo string
}

func Exists(name string) bool {
    if _, err := os.Stat(name); err != nil {
    if os.IsNotExist(err) {
    	return false
	}
    }
    return true
}

func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return nil
}

func IfDirectory(dirname string) bool {
    f, err := os.Open(dirname)
    if err != nil {
        //fmt.Println(err)
	return false
    }
    defer f.Close()
    fi, err := f.Stat()
    if err != nil {
        //fmt.Println(err)
	return false
    }
	switch mode := fi.Mode(); {
    	case mode.IsDir():
	  //fmt.Printf("it's a dir\n")
	case mode.IsRegular():
	  //fmt.Printf("It's a file\n")
	}
	return true
}
