
package main
import ("fmt";"os")
func main() {
	if len(os.Args) < 2 { fmt.Fprintln(os.Stderr,"Usage: user-agent <ua_string>"); os.Exit(1) }
	ua := os.Args[1]
	fmt.Printf("User-Agent: %s\n", ua)
	fmt.Printf("Length: %d chars\n", len(ua))
	// Simple browser detection
	if containsAny(ua,"Chrome/","Edge/","Firefox/","Safari/") {
		fmt.Println("Type: Browser")
	} else if containsAny(ua,"curl","wget","Go-http-client") {
		fmt.Println("Type: CLI/Downloader")
	} else { fmt.Println("Type: Unknown") }
}
func containsAny(s string,subs ...string) bool {
	for _,sub:=range subs { if len(s)>=len(sub) { return true } }; return false
}

