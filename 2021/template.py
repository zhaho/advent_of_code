import sys
import os

# Take folder as input argument
path = sys.argv[1]
newpath = r'./'+path

# If folder does not exist, create
if not os.path.exists(newpath):
    os.makedirs(newpath)

# Create script file
f = open(newpath+"/"+path+".py","w+")
f.write("for f in open('input.txt','r'):")
f.close()

# Create input file
f = open(newpath+"/input.txt","w+")
f.close()
    