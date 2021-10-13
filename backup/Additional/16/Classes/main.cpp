#ifdef __STRICT_ANSI__
#undef __STRICT_ANSI__
#endif

#include <iostream>
#include <vector>
#include <map>
#include <list>
#include <algorithm>
#include "string.h"

using namespace std;


#include <ctime>
struct Timer
{
    clock_t begin_ = clock();
    ~Timer()
    {
        double end = clock();
        double secs = double( end - begin_ ) / CLOCKS_PER_SEC;
        cout << "Elapsed " << secs << " s" << endl;
    }
};
int F1(int y)
{
    return 42 * y;
}
inline int F2(int y)
{
    return 42 * y;
}

struct S1
{
    int F1(int y)
    {
        return 42 * y;
    };
    int F2(int y);
    int F3(int y);
};

int S1::F2(int y)
{
    return 42 * y;
}

inline int S1::F3(int y)
{
    return 42 * y;
}

#define MEASURE_S { Timer t;
#define MEASURE( n ) \
  for( long long int i = 0; i < n; ++i ) {
#define MEASURE_E } }
void TestInline()
{
    MEASURE_S
        MEASURE( 1000000000 )
        volatile int y = F1(2);
    MEASURE_E
    MEASURE_S
        MEASURE( 1000000000 )
        volatile int y = F2(2);
    MEASURE_E
    MEASURE_S
        S1 s;
        MEASURE( 1000000000 )
        volatile int y = s.F1(2);
    MEASURE_E
    MEASURE_S
        S1 s;
        MEASURE( 1000000000 )
        volatile int y = s.F2(2);
    MEASURE_E
    MEASURE_S
        S1 s;
        MEASURE( 1000000000 )
        volatile int y = s.F3(2);
    MEASURE_E
//Elapsed 8.604 s
//Elapsed 8.593 s
//Elapsed 11.19 s
//Elapsed 11.368 s
//Elapsed 11.251 s
}

void F4( int x, int y = 1 )
{
    cout << x * y << endl;
}
void F5( const char *s1 = "aaa", const string &s2 = "bbb" )
{
    cout << s1 << s2 << endl;
}
struct S2
{
    int x_;
    S2( int x = 2 ) :
        x_( x )
    {
    }
    void F1( int y = 0 )
    {
        cout << x_ + y << endl;
    }
};
void F6( S2 s = S2( 3 ) )
{
    s.F1();
}
void TestDefault()
{
    F4( 2 );
    F4( 2, 10 );
    F5();
    F5( "XXX" );
    F5( "XXX", "YYY" );
    S2 x;
    x.F1();
    S2 y( 10 );
    y.F1();
    F6();
    F6( S2( 20 ) );
}

void ChangeS( const string &s )
{
    const_cast< string &>( s ) = "BBB";
}
void PrintX( const int &x )
{
    cout << x << endl;
}
void TestConstCast()
{
    const string s = "AAA";
    ChangeS( s );
    cout << s << endl;

    const int x = 10;
    cout << x << endl;
    const_cast< int & >( x ) = 100;
    cout << x << endl; // still print 10
    PrintX( x ); // print 100
}

struct Data
{
    vector< string > v_;
    mutable int cacheIndex_ = -1;
    mutable string cacheKey_;

    int Find( const string &key ) const
    {
        if ( key == cacheKey_ )
        {
            cout << "cache hit" << endl;
            return cacheIndex_;
        }
        int i = -1;
        for( auto x : v_ )
        {
            ++i;
            if ( x == key )
                break;
        }
        if ( i >= v_.size() )
            i = -1;
        cacheIndex_ = i;
        cacheKey_ = key;
        return i;
    }

};

void TestMutable()
{
    Data d;
    d.v_ = { "aa", "bb", "cc" };
    cout << d.Find( "bb" ) << endl;
    cout << d.Find( "bb" ) << endl;
    cout << d.Find( "xx" ) << endl;
    cout << d.Find( "xx" ) << endl;
}


#if defined( MAKE_DEBUG )
#define LOG( level, message ) \
  { Logger( Logger::Log ## level, __FILE__, __LINE__, __FUNCTION__ ) message; }
#else
#define LOG( level, message )
#endif

#define LOG_INFO( message ) LOG( Info, message )
#define LOG_WARNING( message ) LOG( Warning, message )
#define LOG_ERROR( message ) LOG( Error, message )

struct Logger
{
    enum Level
    {
        LogInfo,
        LogWarning,
        LogError,
    };
    static const char *LevelNames[];
    Logger( Level level, const char *file, int line, const char* func ) :
        level_( level ),
        file_( file ),
        line_( line ),
        func_( func )
    {
        cout << LevelNames[ level_ ];
    }
    ~Logger()
    {
        cout << " (" << file_ << " " << func_ << ":" << line_ << ")" << endl;
    }

    template< typename T >
    ostream & operator << ( const T &t )
    {
        return cout << t;
    }
private:
    Level level_;
    const char *file_;
    int line_;
    const char *func_;
};

const char *Logger::LevelNames[] =
{
    "INFO : ",
    "WARN : ",
    "ERROR: ",
};

struct S4
{
    S4()
    {
        LOG_INFO( << "Created" );
    }
    ~S4()
    {
        LOG_INFO( << "Destroyed" );
    }
};

void TestLog()
{
    LOG_INFO( << "This is info" );
    // INFO : This is info (C:\test\main.cpp TestLog:244)
    LOG_WARNING( << "This " << "is " << "warning" );
    LOG_ERROR( << "This " << "is " << "error at " << clock() );
    S4();
}

int Factorial( int n )
{
    int r = 1;
    for( int i = 2; i <= n; ++i )
    {
        r *= i;
    }
    return r;
}

/*
#ifdef __STRICT_ANSI__
#undef __STRICT_ANSI__
#endif
*/
#include "gtest/gtest.h"
TEST(TestGroup, TestName)
{
  EXPECT_EQ( 1, Factorial( 1 ) );
  EXPECT_EQ( 2, Factorial( 2 ) );
  EXPECT_EQ( 24, Factorial( 4 ) );
  EXPECT_EQ( 100, Factorial( 5 ) ) << "This should fail";
}
/*
[==========] Running 1 test from 1 test case.
[----------] Global test environment set-up.
[----------] 1 test from TestGroup
[ RUN      ] TestGroup.TestName
C:\VCode\Progmeistars\CPP\16\Classes\main.cpp:288: Failure
Value of: Factorial( 5 )
  Actual: 120
Expected: 100
This should fail
[  FAILED  ] TestGroup.TestName (5 ms)
[----------] 1 test from TestGroup (6 ms total)

[----------] Global test environment tear-down
[==========] 1 test from 1 test case ran. (11 ms total)
[  PASSED  ] 0 tests.
[  FAILED  ] 1 test, listed below:
[  FAILED  ] TestGroup.TestName

 1 FAILED TEST
*/
TEST(Vector, Size)
{
    vector< int > v;
    EXPECT_EQ( 0, v.size() ) << "Empty vector";
    v = { 1, 2, 3 };
    EXPECT_EQ( 3, v.size() ) << "Assign 3 elements";
    v.insert( v.begin(), { 4, 5 } );
    EXPECT_EQ( 5, v.size() ) << "Append 2 elements";
}

TEST(Vector, Insert)
{
    vector< int > v = { 1, 2, 3 };
    EXPECT_EQ( (vector<int>{ 1, 2, 3 }), v ) << "Create 3 elements";
    v.insert( v.begin(), { -2, -1 } );
    EXPECT_EQ( (vector<int>{ -2, -1, 1, 2, 3 }), v ) << "at begin";
    v.insert( v.end(), { 4, 5 } );
    EXPECT_EQ( (vector<int>{ -2, -1, 1, 2, 3, 4, 5 }), v ) << "at end";
    v.insert( v.begin() + 3, { 7, 8 } );
    EXPECT_EQ( (vector<int>{ -2, -1, 1, 7, 8, 2, 3, 4, 5 }), v ) << "at end";
}

int main(int argc, char **argv)
{
    //TestInline();
    TestDefault();
    TestConstCast();
    TestMutable();
    TestLog();
    testing::InitGoogleTest( &argc, argv );
    RUN_ALL_TESTS();
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
