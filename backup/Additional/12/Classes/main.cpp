#include <iostream>
#include <vector>
#include <memory>
#include <string.h>

using namespace std;

struct A
{
    int x_;
    virtual ~A(){};
};
struct B : public A
{
};
struct C : public A
{
    void Init()
    {
        // ...
    }
    void Done()
    {
        // ...
    }
};

#include <stdexcept> // out_of_range
#include <typeinfo> // bad_cast
//www.cplusplus.com/reference/exception/exception/
void TestCatch()
{
    try
    {
        int *d = new int[1000000000];
        cout << "You will not see this" << endl;
    }
    catch( const exception &e )
    {
        cout << "std::exception: " << e.what() << endl;
    }
    try
    {
        vector< int > v( 10 );
        v.at(20) = 100;
    }
    catch( const out_of_range &e )
    {
        cout << "Out of Range error: " << e.what() << endl;
    }
    try
    {
        A *b = new B();
        C &c = dynamic_cast< C & >( *b );
    }
    catch( const bad_cast &e )
    {
        cout << "Bad cast error: " << e.what() << endl;
    }
    catch( const exception &e )
    {
        cout << "Other error: " << e.what() << endl;
    }
}

void ThrowInt()
{
    throw 42;
}
void ThrowCString()
{
    throw "Exception!";
}
void ThrowB()
{
    throw B();
}
void ThrowB2()
{
    B b;
    b.x_ = 42;
    throw b;
}
void TestThrow()
{
    try
    {
        ThrowInt();
    }
    catch ( int &x )
    {
        cout << "Catch " << x << endl;
    }
    try
    {
        ThrowCString();
    }
    catch ( const char *x )
    {
        cout << "Catch " << x << endl;
    }
    try
    {
        ThrowB();
    }
    catch ( const A &a )
    {
        cout << "Catch A " << a.x_ << endl;
    }
    try
    {
        ThrowB2();
    }
    catch ( const A &a )
    {
        cout << "Catch A " << a.x_ << endl;
    }
}

 void TestFinal()
 {

    int *p = new int;
    try
    {
        ThrowInt();
        delete p;
    }
    catch(...)
    {
        delete p;
        cout << "something wrong" << endl;
    }

    p = new int;
    try
    {
        try
        {
            ThrowInt();
        }
        catch( const A &a )
        {
            cout << "A" << endl;
        }
        delete p;
    }
    catch(...)
    {
        delete p;
    }

    unique_ptr< int > p2( new int );
    try
    {
        ThrowInt();
    }
    catch( const bad_cast &e )
    {
        cout << "Bad cast error: " << e.what() << endl;
    }
    catch( const exception &e )
    {
        cout << "Other error: " << e.what() << endl;
    }
    catch(...)
    {
        cout << "can't deal with this" << endl;
        // throw; // re-throw, someone else can catch it
    }
 }

 void TestConstructor()
 {
     C c;
     try
     {
        c.Init();
     }
     catch(...)
     {
         // ....
     }
     cout << c.x_ << endl;
     try
     {
        c.Done();
     }
     catch(...)
     {
         // ....
     }
 }

void F1() throw( int, const char * )
//void F1() throw( const char * )
{
    ThrowInt();
}

void F2() throw()
{
    ThrowInt();
}

 void TestSpec()
 {
    try
    {
        F1();
    }
    catch( int )
    {
        cout << "int" << endl;
    }
    try
    {
        //F2(); // will abort program
    }
    catch( ... )
    {
        cout << "something" << endl;
    }
 }

class Backup
{
private:
    char *data_;
    unsigned int size_;
    vector< char > copy_;

public:
    Backup( void *data, unsigned int dataSize )
    {
        data_ = static_cast< char *>( data );
        size_ = dataSize;
        copy_.assign( data_, data_ + size_ );
    }
    ~Backup()
    {
        if ( data_ )
        {
            memcpy( data_, &copy_[0], size_ );
        }
    }
    void Release()
    {
        data_ = 0;
    }
};

void TestScopedUsed()
{
    char s[100];
    strcpy( s, "123456" );
    cout << s << endl;
    try
    {
        Backup b( s, sizeof( s ) );
        s[3] = 'x';
        ThrowInt();
        b.Release();
    }
    catch( int )
    {
    }
    cout << s << endl;
    try
    {
        Backup b( s, sizeof( s ) );
        s[3] = 'y';
        //ThrowInt();
        b.Release();
    }
    catch( int )
    {
    }
    cout << s << endl;
}

enum Number
{
    ZERO,
    ONE,
    TWO,
    TEN = 10,
    MINUS_ONE = -1,
};

enum Options
{
    OPT1,
    OPT2,
    OPT3,
    OPT4,
    OPT_SIZE,
};

struct Data
{
    enum
    {
        MAX_SIZE = 100,
        MIN_SIZE = 10,
    };
};

void FEnum( Number n )
{
    cout << n << endl;
}

void TestEnum()
{
    cout << TEN << endl;
    Number x = TWO;
    cout << x << endl;
    // x = 3; // Invalid conversation
    switch( x )
    {
        case ZERO:
            //...
            break;
        case ONE:
        case TWO:
            //...
            break;
        // without default will be warning
        // because TEN, MINUS_ONE is not handled
        default:
            //...
            break;
    }
    cout << Data::MAX_SIZE << endl;
    FEnum( MINUS_ONE );
    int a = -1;
    //FEnum( a );
    FEnum( Number( a ) );
    FEnum( Number( 123 ) );
    int b = TEN; // all OK

    int opt = 4;
    if ( opt > OPT_SIZE )
    {
        cout << "Invalid option " << x << endl;
    }

}



int main()
{
    TestCatch();
    TestThrow();
    TestFinal();
    TestConstructor();
    TestSpec();
    TestScopedUsed();
    TestEnum();
    return 0;
}
