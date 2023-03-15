set msg [gets stdin]

set mode [lindex $argv 0]
if { $mode eq "-d" } {
	set text [binary decode base64 $msg]
	puts $text
} elseif { $mode eq "-e" } {
	set text [binary encode base64 $msg]
	puts $text
}