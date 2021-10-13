#include <iostream>
#include <vector>
#include <map>
#include <list>
#include <algorithm>
#include "string.h"

#include "SomeData.h"

using namespace std;

void TestLib()
{
    SomeData d1(10);
    d1.Print();
    SomeData d2 = CreateSomeData();
    d2.Print();
}

extern "C" void Foo( int x )
{
    cout << "It is foo " << x << endl;
}

extern "C"
{

void Bar1( int x )
{
    cout << "It is bar1 " << x << endl;
}

void Bar2( int x )
{
    cout << "It is bar2 " << x << endl;
}

}

void TestExternC()
{
    Foo( 10 );
    Bar1( 20 );
    Bar2( 30 );
}

struct Obj
{
    char *big;
    Obj()
    {
        cout << "Created[def] at " << (unsigned int)this << endl;
        big = new char[1000];
    }
    Obj( const Obj &other )
    {
        cout << "Created[copy] at " << (unsigned int)this << endl;
        big = new char[1000];
        memcpy( big, other.big, 1000 );
    }

    Obj( Obj &&other )
    {
        cout << "Move at " << (unsigned int)this << endl;
        big = other.big;
        other.big = nullptr;
    }

    ~Obj()
    {
        cout << "Destroyed at " << (unsigned int)this << endl;
        delete[] big;
    }
};

#include <utility>
void TestMove()
{
    {
        Obj x;
        vector< Obj > v;
        v.reserve(2);
        v.push_back( x );
        v.push_back( Obj() );
        //Output if no Move constructor exists:
        //Created[def] at 2686696
        //Created[copy] at 9056088
        //Created[def] at 2686708
        //Created[copy] at 9056092
        //Destroyed at 2686708
        //Destroyed at 9056088
        //Destroyed at 9056092
        //Destroyed at 2686696
    }
    string s1 = "SOME TEXT";
    string s2 = move( s1 );
    cout << s1 << "," << s2 << endl;
    {
        Obj x;
        vector< Obj > v;
        v.reserve(2);
        v.push_back( move( x ) );
        v.push_back( Obj() );
        //Created[def] at 2686680
        //Move at 7679832
        //Created[def] at 2686716
        //Move at 7679836
        //Destroyed at 2686716
        //Destroyed at 7679832
        //Destroyed at 7679836
        //Destroyed at 2686680
    }
}

struct Base
{
    int F1()
    {
        return 1;
    }
};

struct Child : public Base
{
    int F1()
    {
        return 2;
    }
    void Process()
    {
        cout << F1() << endl;
        cout << Base::F1() << endl;
    }
};


void TestBaseClassCall()
{
    Child c;
    c.Process();
}

struct Spacer
{
    string s_;
    Spacer( int i ) :
        s_( i, '_' )
    {
    }
    const string & str() const
    {
        return s_;
    }
};

void Print( int x, int y, const Spacer &sp )
{
    cout << x << sp.str() << y << endl;
}

void TestExplicit()
{
    Spacer s(4);
    cout << s.str() << endl;
    // This is not obvious
    // Print( 1, 2, 3 ); // print '1___2'
    Print( 1, 2, Spacer( 3 ) );
    // This is not obvious
    // s = 2;
    // s = 'a';
    //en.wikibooks.org/wiki/More_C%2B%2B_Idioms/Safe_bool
}

int main()
{
    TestLib();
    TestExternC();
    TestMove();
    TestBaseClassCall();
    TestExplicit();
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
