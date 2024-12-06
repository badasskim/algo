inputedNum = int(input())
countNum = 0
trippleSix = 666

while True:
    if '666' in str(trippleSix):
        countNum += 1 
    
    if countNum == inputedNum:  
        break
    
    trippleSix += 1  
    
print(trippleSix)