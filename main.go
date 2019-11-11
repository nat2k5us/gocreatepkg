package main
import (
	 "fmt"
	 "ellipse"
	 
)
func main() {
	if gopath := cfg.BuildContext.GOPATH; filepath.Clean(gopath) == filepath.Clean(runtime.GOROOT()) { 
		fmt.Fprintf(os.Stderr, "warning: GOPATH set to GOROOT (%s) has no effect\n", gopath) 
	} else { 
	 //Initalise the Init function with value of A,B
	 e := ellipse.Init{
		9, 2,
	   }
	   //this will give answer as 0.9749960430435691
	   fmt.Println(e.GetEccentricity())
	}
}

