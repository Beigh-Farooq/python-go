class Node:
    __slots__='element','next'
    def __init__(self,element,next):
        self.element=element 
        self.next=next 
class LinkedList:
    def __init__(self):
        self.head=None 
        self.tail=None  
        self.size=0
    def __len__(self):
        return self.size 
    def isempty(self):
        return self.size==0
    
    def addlast(self,e):
        newest=Node(e,None)
        if self.isempty():
            self.head=newest
        else:
            self.tail.next=newest
        self.tail=newest
        self.size+=1
    def removeduplicate(self):
        p=self.head
        while p:
            if p.element==p.next.element:
                p.next=p.next.next
        p=p.next
        return self.head
        
    def display(self):
        p=self.head
        while p:
            print(p.element,end='-->')
            p=p.next
        print()
L=LinkedList()
L.addlast(1)
L.addlast(2)
L.addlast(2)
L.addlast(3)
L.display()
print(L.removeduplicate())
