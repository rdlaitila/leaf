import sys, os, shutil, subprocess, shlex
from subprocess import Popen, PIPE

#Base path to our script. Move up two parent directories to get into the repo base folder
BASE_PATH = os.path.dirname(os.path.dirname(os.path.dirname(os.path.realpath(__file__))))

#Path to our pkg directory
PKG_PATH = BASE_PATH + "/pkg"

#Path to our release build directory
RLS_PATH = BASE_PATH + "/release"

#Path to our temporary directory
TMP_PATH = BASE_PATH + "/tmp"

#Path to our src directory
SRC_PATH = BASE_PATH + "/src"

#Paths to search for C source files
LIBRARY_PATH = BASE_PATH + "/src/_lua"

#Path to search for C source files. This is just a alias to LIBRARY_PATH 
LD_LIBRARY_PATH = LIBRARY_PATH

#Cgo LDFlags used in go compilation of C programs
CGO_LDFLAGS =  [
    '-L'+BASE_PATH+'/src/_lua/ ',
    '-lm ',
    '-ldl '
]

#Cgo CFLAGS used in go compilation of C programs
CGO_CFLAGS = [
    "-I"+BASE_PATH+"/src/_lua/ "
]


#
#.SYNOPSIS
#   main is the entry point to the script. Called from the very bottom of the script to ensure
#   all function declarations have been parsed.
#
def main():
    #output some of our constants for quick debugging
    log("BASE_PATH: " + BASE_PATH)
    log("TMP_PATH: " + TMP_PATH)
    log("RLS_PATH: " + RLS_PATH)
    log("LIBRARY_PATH: " + LIBRARY_PATH)
    log("LD_LIBRARY_PATH: " + LD_LIBRARY_PATH)
    log("CGO_LDFLAGS: " + " ".join(CGO_LDFLAGS))
    log("CGO_CFLAGS: " + " ".join(CGO_CFLAGS))

    if sys.argv[1] == "--build":
        log("Starting Project Build")
        build()
    elif sys.argv[1] == "--test":
        log("Starting Project Tests")
        test()
    elif sys.argv[1] == "--clean":
        log("Starting Project Clean")
        clean()
    

#
#.SYNOPSIS
#    build builds the software
#    
def build():
    #Remove .\pkg and .\tmp directories if they exist
    if os.path.exists(TMP_PATH):
        shutil.rmtree(TMP_PATH)
    if os.path.exists(PKG_PATH):
        shutil.rmtree(TMP_PATH)
        
    #Create our temporary working directory
    os.makedirs(TMP_PATH)
        
    #Compile leaf programs
    os.chdir(BASE_PATH)
    os.environ['GOPATH'] = BASE_PATH
    os.environ['CGO_LDFLAGS'] =  " ".join(CGO_LDFLAGS)
    os.environ["CGO_CFLAGS"] = " ".join(CGO_CFLAGS)
    
    #Compile leaf
    cmd = ["go", "build", "src/_cmd/leaf.go"]
    log("Invoking Command: " + str(cmd))
    exit_code = execute(cmd)
    if exit_code != 0:
        sys.exit(exit_code)
    shutil.move(BASE_PATH+"/leaf", TMP_PATH+"/leaf")
    
    #Compile leaf51
    os.environ['CGO_LDFLAGS'] = " ".join(CGO_LDFLAGS)
    cmd = ["go", "build", "-x", "src/_cmd/leaf51.go"]
    log("Invoking Command: " + str(cmd))
    exit_code = execute(cmd)
    if exit_code != 0:
        sys.exit(exit_code)
    shutil.move(BASE_PATH+"/leaf51", TMP_PATH+"/leaf51")
    
    #Compile leaf52
    os.environ['CGO_LDFLAGS'] = " ".join(CGO_LDFLAGS)
    cmd = ["go", "build", "-x", "src/_cmd/leaf52.go"]
    log("Invoking Command: " + str(cmd))
    exit_code = execute(cmd)
    if exit_code != 0:
        sys.exit(exit_code)
    shutil.move(BASE_PATH+"/leaf52", TMP_PATH+"/leaf52")
    
    #Compile leaf53
    os.environ['CGO_LDFLAGS'] = " ".join(CGO_LDFLAGS)
    cmd = ["go", "build", "-x", "src/_cmd/leaf53.go"]
    log("Invoking Command: " + str(cmd))
    exit_code = execute(cmd)
    if exit_code != 0:
        sys.exit(exit_code)
    shutil.move(BASE_PATH+"/leaf53", TMP_PATH+"/leaf53")
    

#
#.SYNOPSIS
#   test runs all available test applications
#    
def test():
    print "test"

#
#.SYNOPSIS
#    cleans the build
#    
def clean():
    print "clean"

#
#.SYNOPSIS
#    log simply outputs a log line to shell
#
def log(msg, lvl="INFO"):
    print lvl + ": " + msg

#
#.SYNOPSIS
#   executes a command
#
def execute(command):    
    popen = subprocess.Popen(command, stdout=subprocess.PIPE)
    lines_iterator = iter(popen.stdout.readline, b"")
    for line in lines_iterator:
        print(line) # yield line
    return popen.wait()
    
#
# Call main
#
main()