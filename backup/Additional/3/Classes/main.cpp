#include <iostream>
#include <cmath>

using namespace std;

class Parent
{
public:
    void M1()
    {
        cout << "Parent::M1" << endl;
    }
    virtual void M2()
    {
        cout << "Parent::M2" << endl;
    }
    void M3()
    {
        cout << "Parent::M3" << endl;
        M4();
    }
    virtual void M4()
    {
        cout << "Parent::M4" << endl;
    }
};

class Child : public Parent
{
public:
    void M1()
    {
        cout << "Child::M1" << endl;
    }
    virtual void M2()
    {
        cout << "Child::M2" << endl;
    }
    virtual void M4()
    {
        cout << "Child::M4" << endl;
    }
};

void Caller1( Parent *p )
{
    p->M1();
    p->M2();
}

void Caller2( Parent *p )
{
    p->M3();
}

void TestVirtual()
{
    Parent p;
    Caller1( &p ); // Parent::M1\nParent::M2
    Child c;
    Caller1( &c ); // Parent::M1\nChild::M2

    Caller2( &p );
    Caller2( &c );
}

typedef Parent * PParent;
void TestMemory()
{
    {
        int *p = new int;
        *p = 100;
        cout << *p << endl;
        delete p;
        p = 0;

        Parent *base = new Child();
        base->M1();
        base->M2();
        delete base;
        base = 0;
    }
    {
        int *p = new int[10];
        for( int i = 0; i < 10; ++i )
        {
            p[i] = i * i;
        }
        delete[] p;
        p = 0;

        // Create using default constructor
        Parent *bases = new Parent[5];
        for( int i = 0; i < 5; ++i )
        {
            bases[i].M1();
            bases[i].M2();
        }
        delete[] bases;
        bases = 0;

        // 10 pointers to Parent, no memory allocation
        // Parent *bases2[5];
        PParent bases2[5];
        cout << "Virtual test" << endl;
        for( int i = 0; i < 5; ++i )
        {
            if ( ( i % 2 ) == 0 )
            {
                bases2[i] = new Parent();
            }
            else
            {
                bases2[i] = new Child();
            }
            bases2[i]->M1();
            bases2[i]->M2();
        }

        for( int i = 0; i < 5; ++i )
        {
            delete bases2[i];
            bases2[i] = 0;
        }
    }
}

class ClassA
{
public:
    ClassA()
    {
        cout << "ClassA::ClassA" << endl;
    }
    virtual ~ClassA()
    {
        cout << "ClassA::~ClassA" << endl;
    }

};

class ClassB : public ClassA
{
public:
    int *p;
    ClassB() :
        ClassA()
    {
        p = new int[10];
        cout << "ClassB::ClassB, new 10 ints" << endl;
    }
    virtual ~ClassB()
    {
        delete p;
        cout << "ClassB::~ClassB, free 10 ints" << endl;
    }

};

void TestDestructors()
{
    {
        ClassA a;
        ClassB b;
    }
    {
        ClassA *base = new ClassB();
        // If destructor is not virtual then ~ClassA will be called
        // and memory is not freed
        delete base;
    }
}

void ChangeIntPointer( int *x )
{
    *x = 100;
}

void ChangeIntRef( int &x )
{
    x = 200;
}

void Caller3( Parent &p )
{
    p.M1();
    p.M2();
}

void TestReferences()
{
    int x;
    ChangeIntPointer( &x );
    cout << x << endl;
    ChangeIntRef( x );
    cout << x << endl;

    int *pX = &x;
    *pX = 300;
    cout << x << " " << *pX << endl;

    int &refX = x;
    refX = 400;
    cout << x << " " << refX << endl;

    Parent *p = new Child();
    Caller3( *p );
    delete p;
    p = 0;

    //Caller3( *p ); // runtime error when call M2(because virtual)
    int *pA = 0;
    int &rA = *pA;
    //cout << rA << endl; // runtime error
}

int main()
{
    TestVirtual();
    TestMemory();
    TestDestructors();
    TestReferences();

    return 0;
}
