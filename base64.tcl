package require base64

proc split {str size} {
    set result [list]
    for {set i 0} {$i < [string length $str]} {incr i $size} {
        lappend result [string range $str $i [expr {$i + $size - 1}]]
    }
    return $result
}

set col 64
set dec 0
set pad 0

set mode [lindex $argv 0]

if { $mode eq "-d" } {
    set dec 1
} elseif { $mode eq "-e" } {
    set dec 0
}

set inputData ""
if {[eof stdin] == 0} {
    set inputData [read stdin]
} elseif {[llength $argv] > 1} {
    set inputFile [lindex $argv 1]

    set fd [open $inputFile r]
    set inputData [read $fd]
    close $fd
}

if { $dec } {
    # Se estiver decodificando, remover as quebras de linha apenas do texto codificado
    set inputData [string map {\r "" \n ""} $inputData]
}

if { $col == 0 } {
    if { !$dec && !$pad } {
        set sEnc [::base64::encode $inputData]
        puts $sEnc
    } elseif { $dec && !$pad } {
        set decoder [::base64::decode $inputData]
        puts $decoder
    }

    if { !$dec && $pad } {
        set sEnc [::base64::encode -nopad $inputData]
        puts $sEnc
    } elseif { $dec && $pad } {
        set decoder [::base64::decode -nopad $inputData]
        puts $decoder
    }
} else {
    if { !$dec && !$pad } {
        set sEnc [::base64::encode $inputData]
        set sEnc [string map { "\n" "" } $sEnc] ;# Remover quebras de linha do texto codificado
        set chunks [split $sEnc $col]
        foreach chunk $chunks {
            puts -nonewline "$chunk\n"
        }
    } elseif { $dec && !$pad } {
        set decoder [::base64::decode $inputData]
        puts $decoder
    }

    if { !$dec && $pad } {
        set sEnc [::base64::encode -nopad $inputData]
        set sEnc [string map { "\n" "" } $sEnc] ;# Remover quebras de linha do texto codificado
        set chunks [split $sEnc $col]
        foreach chunk $chunks {
            puts -nonewline "$chunk\n"
        }
    } elseif { $dec && $pad } {
        set decoder [::base64::decode -nopad $inputData]
        puts $decoder
    }
}
