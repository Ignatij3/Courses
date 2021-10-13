#include <iostream>
#include <vector>
#include <map>
#include <list>
#include <algorithm>
#include "string.h"

using namespace std;


constexpr int F1( int x )
{
    return 2 * x;
}
const int X1 = F1( 10 );

constexpr int valueSize()
{
    return 5 + F1(2);
}

void TestConstExpr()
{
    int value[ valueSize() ];
    cout << sizeof( value ) << endl;
}

vector< int > F2()
{
    return { 1, 2, 3 };
}

struct Data
{
    vector< int > d;
    Data( initializer_list<int> init )
    {
        for( const int *i = init.begin(); i != init.end(); ++i )
        {
            d.push_back( *i );
        }
    }

    void Append( initializer_list<int> l )
    {
        d.insert( d.end(), l.begin(), l.end() );
    }
};

void TestInitializerList()
{
    vector< int > v = { 1, 2, 3, 4 };
    cout << v.size() <<  " " << v.back() << endl;
    map< string, int > m = { {"word1", 1}, {"word2", 2} };
    cout << m.size() << m.begin()->first << ":" << m.begin()->second << endl;

    vector<int> v2 = F2();
    cout << v2[0] << v2[1] << v2[3] << endl;

    Data d = { 1, 2, 3 };
    d.Append( { 4, 5, 6 } );
    cout << d.d.size() << endl;
    vector< Data > d2 = { Data{1,2,3}, Data{1, 2} };
    cout << d2.size() << d2[0].d.size() << endl;

    vector< pair<string, string> > people =
        { {"Alex", "Eve"}, {"Bob", "Ada"} };
    cout << people[0].first << " and " << people[0].second << endl;
    vector< vector<int> > v3 = { { 1, 2, 3 }, {4, 5 } };
    cout << v3.size() << v3[0].size() << v3[1].size() << endl;

    struct S1 {
        int f1;
        float f2;
        string f3;
    };
    // Uniform initialization
    S1 s1 = { 1, 2.0, "test" };
    cout << s1.f3 << endl;
    S1 s2 = { 1, 2.0 };
    cout << s2.f3 << endl;

}

int F3( double d, const string &s )
{
    cout << d << s << endl;
}

void TestAuto()
{
    int x = 3;
    auto x2 = 3;
    cout << typeid( x ).name() << " " << typeid( x2 ).name() << endl;
    list< int > l1 = {1, 2, 3};
    for( auto it = l1.begin(); it != l1.end(); ++it )
    {
        *it *= 2;
    }
    for( auto it = l1.cbegin(); it != l1.cend(); ++it )
    {
        cout << *it << endl;
    }

    //function< int( double d, const string &s ) > f =
    // bind( &F3, placeholders::_1, placeholders::_2 );
    auto f = bind( &F3, placeholders::_1, placeholders::_2 );
    f( 70, "aa" );

    int i0 = 1;
    decltype( i0 ) i1 = 3;
    decltype( F3(1, "") ) i2 = 4;
    cout << i1 << " " << i2 << endl;

    vector<int> v = { 1, 2, 3 };
    decltype( v[0] ) i3 = v[0];
    i3 = 4;
    cout << v[0] << endl;

    decltype( (i1) ) i4 = i1;
    i4 = 5;
    cout << i1 << endl;
}

void TestRangeFor()
{
    list< int > l1 = {1, 2, 3};
    for( auto &x : l1 )
    {
        x *= 2;
    }
    for( const auto &x : l1 )
    {
        cout << typeid(x).name() << x << endl;
    }
}

struct Data3
{
    vector<int> v={1, 2, 3};
    int total = 0;
    void Foo()
    {
        for_each( v.begin(), v.end(), [this] (int x) {
          this->total += x;
        });
    }
};

void TestLambda()
{
    vector<int> v{ 10, -5, 2, -6};
    sort( v.begin(), v.end(), [](int a, int b)
        { return abs(a)<abs(b); } );
    for( auto &x : v )
    {
        cout << x << endl;
    }
    int total = 0;
    for_each( v.begin(), v.end(), [&total] (int x) {
        total += x;
    });
    cout << total << endl;
    auto f = find_if( v.begin(), v.end(), [=](int x)
        { return x > total; } );
    cout << *f << " found" << endl;

    Data3 d;
    d.Foo();
    cout << d.total << endl;
}

struct Base {
    virtual void Func(float);
};

struct Derived : Base {
    // this will give error, because Func(int)
    // create second virtual function
    // virtual void Func(int) override;

    virtual void Func(float) override;
};

class SomeType  {
    int number;

public:
    SomeType(int new_number) : number(new_number) {}
    SomeType() : SomeType(42) {}
};

struct NonCopyable {
    NonCopyable() = default;
    NonCopyable(const NonCopyable&) = delete;
    NonCopyable & operator=(const NonCopyable&) = delete;
    int Foo() { return 42; }
};
class X : public NonCopyable
{
    int Foo() = delete;
};

struct Base1 final { };
//struct Derived1 : Base1 { };
struct Base2 {
    virtual void f() final;
};

struct Derived2 : Base2 {
    //virtual void f();
};

void TestClass()
{
    char *pc = nullptr;     // OK
    int  *pi = nullptr;     // OK
    bool   b = nullptr;     // OK. b is false.
    // int    i = nullptr;     // error

    X x;
    X x2;
    //x = x2;
    //x.Foo();
}

int main()
{
    TestConstExpr();
    TestInitializerList();
    TestAuto();
    TestRangeFor();
    TestLambda();
    return 0;
}



/*
#include <regex>
void TestRegex()
{

    if (std::regex_match ("subject", std::regex("(sub)(.*)") ))
        std::cout << "string literal matched\n";

    // Simple regular expression matching
    std::string fnames[] = {"foo.txt", "bar.txt", "baz.dat", "zoidberg"};
    std::regex txt_regex("[a-z]+\\.txt", regex_constants::basic);

    for (const auto &fname : fnames) {
        std::cout << fname << ": " << std::regex_match(fname, txt_regex) << '\n';
    }

    // Default to ECMAScript, this one use extended(POSIX regular expression)
    // http://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap09.html#tag_09_04
    string s = "ABCxyz";
    regex reLetters( "[A-Z]BC.*", regex_constants::extended );
    if ( regex_match( s, reLetters ) )
    {
        cout << "Match found" << endl;
    }
    smatch match;
    regex_match( s, match, reLetters );
    cout << match.str(1) << endl;


}

*/

/*
#include <thread>
#include <chrono>



void TestThread()
{
    class Executor
    {
        void operator()()
        {
            cout << "Thread start" << endl;
            //this_thread::sleep_for( chrono::milliseconds( 1000 ) );
            cout << "Thread end" << endl;
        }
    };

    Executor ex;
    cout << "Create thread" << endl;
    thread t2( ex );
    cout << "Thread terminated" << endl;
}

*/
