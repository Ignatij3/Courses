#include <iostream>
#include <cmath>

using namespace std;


class Figure
{
public:
    virtual ~Figure(){};
    virtual void Draw() = 0;
};

class Square : public Figure
{
public:
    virtual void Draw()
    {
        cout << "Square draw" << endl;
    }
};

void Draw( Figure &f )
{
    f.Draw();
}

void TestAbstract()
{
    // Can't, Figure is abstract class
    //Figure f;
    Square sq;
    sq.Draw();
    Figure &f = sq;
    f.Draw();
    Draw( sq );
}

class IPrintable
{
public:
    virtual ~IPrintable(){};
    virtual void Print() = 0;
};

class Book : public IPrintable
{
public:
    virtual void Print()
    {
        cout << "Printing book content ..." << endl;
    }
};
class Circle : public Figure, public IPrintable
{
public:
    virtual void Draw()
    {
        cout << "Square draw" << endl;
    }

    virtual void Print()
    {
        cout << "Printing circle content ..." << endl;
    }
};

void PrintSomething( IPrintable &p )
{
    p.Print();
}

void TestInterface()
{
    Book b;
    b.Print();
    IPrintable &p = b;
    p.Print();
    PrintSomething( b );
    Circle c;
    PrintSomething( c );
}

class Person
{
public:
    void GetName()
    {
        cout << "I have name" << endl;
    }
};
class Worker
{
public:
    void GetSalary()
    {
        cout << "I have salary" << endl;
    }
};
class Developer : public Person, public Worker
{
};

class A
{
public:
   int foo()
   {
       return 1;
   }
};
class B: public virtual A
{
};
class C : public virtual A
{
};
class D : public B, public C
{
};

void TestMultipleInheritance()
{
    Developer dev;
    dev.GetName();
    dev.GetSalary();

    D d;
    d.foo();
}

#include "Child.h"
void TestCpp()
{
    Parent p;
    p.M1();
    p.M2();

    Parent *p2 = new Child();
    p2->M1();
    p2->M2();
    delete p2;
}

int main()
{
    TestAbstract();
    TestInterface();
    TestMultipleInheritance();
    TestCpp();

    return 0;
}
