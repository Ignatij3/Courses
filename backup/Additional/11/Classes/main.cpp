#include <iostream>
#include <vector>

using namespace std;

#include <typeinfo>
class A
{
public:
    // Without error when dynamic cast
    // source type is not polymorphic
    virtual ~A(){};

    const char *AId()
    {
        return "A";
    }
};
class B : public A
{
    int x;
public:
    const char *BId()
    {
        return "B";
    }
};
class C : public B
{
public:
    const char *CId()
    {
        return "C";
    }
};
class D : public A
{
public:
    const char *DId()
    {
        return "D";
    }
};

typedef int myint;
void TestTypeId()
{
    A a;
    cout << ( typeid( A ) == typeid( a ) ) << endl;
    B b;
    cout << ( typeid( A ) == typeid( b ) ) << endl;
    myint i;
    cout << ( typeid( int ) == typeid( i ) ) << endl;
    A *p = &b;
    cout << ( typeid( B ) == typeid( *p ) ) << endl;
}

void RefStaticCast( B &b )
{
    cout << static_cast< C & >( b ).CId() << endl;
}

void TestStaticCast()
{
    int i = 10;
    double x = 5.5;
    cout << static_cast< int >( x ) << endl; // 5
    cout << static_cast< double >( i ) / 3 << endl;
    //cout << ( char *)&x << endl; // will compile
    //cout << static_cast< char * >( &x ) << endl; // will not

    A *b = new B();
    B *b2 = new B();
    A *c = new C();
    A *d = new D();

    // This is downcasting, it is not always safe
    B *pb = static_cast< B * >( b );
    cout << pb->BId() << endl;
    pb = static_cast< B * >( c );
    cout << pb->BId() << endl;
    C *pc = static_cast< C * >( c );
    cout << pc->CId() << endl;
    //pb = static_cast< D * >( b2 ); // compile error

    // THIS WILL COMPILE, BUT THIS IS ERROR CODE
    // pb = static_cast< B * >( d );
    // cout << pb->BId() << endl;

    C c2;
    D d2;
    RefStaticCast( c2 );
    //RefStaticCast( *b2 );
    //RefStaticCast( d2 );
    delete b; delete b2; delete c; delete d;
}

void PointerDynCast( A *a )
{
    B *b = dynamic_cast< B * >( a );
    if ( b )
    {
        cout << b->BId() << endl;
    }
    C *c = dynamic_cast< C * >( a );
    if ( c )
    {
        cout << c->CId() << endl;
    }
    D *d = dynamic_cast< D * >( a );
    if ( d )
    {
        cout << d->DId() << endl;
    }
}

void RefDynCast( A &a )
{
    B &b = dynamic_cast< B & >( a );
    cout << b.BId() << endl;
}

void TestDynamicCast()
{
    A *a = new A();
    A *b = new B();
    A *c = new C();
    A *d = new D();
    PointerDynCast( a );
    PointerDynCast( b );
    PointerDynCast( c );
    PointerDynCast( d );
    delete a; delete b; delete c; delete d;

    A aa;
    B bb;
    C cc;
    D dd;
    //RefDynCast( aa );
    RefDynCast( bb );
    RefDynCast( cc );
    //RefDynCast( dd );
}

struct BigData
{
    int d[100000];
};

#include <memory>
void DataOut1( BigData *x )
{
    cout << x->d[100] << endl;
}
void DataOut2( BigData &x )
{
    cout << x.d[100] << endl;
}
void TestAutoPtr()
{
    auto_ptr< BigData > a( new BigData );
    a->d[100] = 100;
    DataOut1( a.get() );
    DataOut2( *a );

    BigData *p = new BigData;
    auto_ptr< BigData > a2( p );
    delete p;
    a2.release();

    auto_ptr< A > a3( new C );
    cout << a3->AId() << endl;
    cout << static_cast< C * >( a3.get() )->CId() << endl;
    cout << static_cast< C & >( *a3 ).CId() << endl;
}

struct Figure
{
    Figure()
    {
        cout << "created at " << ( unsigned int)this << endl;
    }
    virtual ~Figure()
    {
        cout << "destroyed at " << ( unsigned int)this << endl;
    }
    virtual void Draw() = 0;
};
struct Circle : public Figure
{
    virtual void Draw()
    {
        cout << "circle draw at " << ( unsigned int)this << endl;
    }
};
struct Square : public Figure
{
    virtual void Draw()
    {
        cout << "square draw at " << ( unsigned int)this << endl;
    }
};

void TestSharedPtr()
{
    cout << "----" << endl;
    typedef shared_ptr< Figure > SPFigure;
    vector< SPFigure > v;
    for( unsigned int i = 0; i < 2; i++ )
    {
        v.push_back( SPFigure( new Square ) );
        v.push_back( SPFigure( new Circle ) );
    }
    for( unsigned int i = 0; i < v.size(); i++ )
    {
        v[i]->Draw();
    }
    vector< SPFigure > v2;
    v2.push_back( v[0] );
    v2.push_back( v[1] );
    for( unsigned int i = 0; i < v2.size(); i++ )
    {
        v2[i]->Draw();
    }
}

void RCastC( void * p )
{
    C * c = reinterpret_cast< C * >( p );
    cout << c->CId() << endl;
}

void TestReinterCast()
{
    int i;
    const char *s = "1234567";
    i = (int)s;
    cout << (char *)i << endl;
    i = reinterpret_cast< int >( s );
    cout << reinterpret_cast< char *>( i ) << endl;

    unique_ptr< C > c( new C );
    RCastC( c.get() );
    unique_ptr< D > d( new D );
    // RCastC( d.get() ); // this is compile but bad
}

int main()
{
    TestTypeId();
    TestStaticCast();
    TestDynamicCast();
    TestAutoPtr();
    TestSharedPtr();
    TestReinterCast();
    return 0;
}
