#include <iostream>
#include <cmath>
#include <cstdio>
using namespace std;

void Foo()
{
    cout << "global" << endl;
}

namespace space1
{

void Foo()
{
    cout << "spacel" << endl;
}

}

namespace space2
{

void Foo()
{
    cout << "space2" << endl;
}

}

// in some.h
namespace space3
{
struct SomeClass
{
    SomeClass();
};
}
// in some.cpp
namespace space3
{
SomeClass::SomeClass()
{
}
}
using space3::SomeClass;

namespace space4
{
double sqrt(double x)
{
    return x;
}
double sin(double x)
{
    return x;
}
}
using namespace space4;

void TestNamespaces()
{
    Foo();
    space1::Foo();
    space2::Foo();
    space3::SomeClass c;
    SomeClass c2;
    cout << space4::sqrt( 4.0 ) << endl;
    cout << ::sqrt( 4.0 ) << endl;
}

void TestString()
{
    string s = "abc";
    cout << s << endl;
    s += "def";
    printf( "as c-string %s\n", s.c_str());
    const char *c = "12345";
    if ( s != c )
    {
        cout << s << " not equal " << c << endl;
    }

    for( int i = 0; i < s.length(); ++i )
    {
        cout << s[i];
        s[i] = toupper(s[i]);
    }
    cout << endl << s << endl;

    string s1 = "123456";
    cout << "found 56 at " <<  s1.find( "56" ) << endl;
    size_t pos = s1.find( "abc" );
    if ( pos == string::npos )
    {
        cout << "abc not found " << endl;
    }
    getline( cin, s );
    cout << "entered = " << s << endl;
    cin >> s; // till space
    cout << "entered = " << s << endl;
}
/* Most usefull :
constructors
assign - Assign content to string
append - Append to string
insert - Insert into string
erase - Erase characters from string
replace - Replace portion of string
substr - Generate substring
copy - Copy sequence of characters from string to C-string
compare - Compare strings
find - Find content in string
rfind - Find last occurrence of content in string

For chars <ctype>:
is*( like isdigit, isalpha, isspace ...
toupper, tolower
*/

#include <iomanip>
void TestMan()
{
    cout << setfill('_') << setw(6) << "six" << "none" << endl;

    double x = 123.45678;
    // default
    cout << x << endl;
    // Maximum number of digits to display
    cout << setprecision( 4 ) << x << endl;
    // Two digits after .
    cout << fixed << setprecision( 2 ) << x << endl;
    // Output in oct/dec/hex for integers
    cout << setbase( 16 ) << 32 << endl;

    // Note: setprecision and setbase have constant effect
    cout << 567.8 << " " << 16 << endl;

    int saveP = cout.precision();
    cout << setprecision(4) << 123.4567 << endl;
    cout.precision( saveP );
    cout << 123.4567 << endl;
}

#include <sstream>
void TestStrNum()
{
    string s = "123.45";
    istringstream is( s );
    double x;
    is >> x;
    cout << x << endl;

    x = 123.45;
    ostringstream os;
    os << fixed << setprecision(1) << x;
    s = os.str();
    cout << s << endl;
}

class Data
{
private:
    int x_;
    string s_;
public:
    Data( int x ) :
        x_( x ),
        s_()
    {

    }
	Data( const Data &other ) :
		x_( other.x_ ),
		s_( other.s_ )
	{
	}
	static const int MAX_SIZE;
	int GetX() const
	{
	    return x_;
	}
	void SetX( int x )
	{
	    x_ = x;
	}
	const string &GetStr() const
	{
	    return s_;
	}
	void SetStr( const string &s )
	{
	    s_ = s;
	}
};
const int Data::MAX_SIZE = 10;

void CallOnConst( const Data &d )
{
    cout << d.GetX() << d.GetStr() << endl;
    // d.SetX( 11 ); // error: passing 'const ...
}

void TestConst()
{
    Data x( 0x10 );
    Data y = x; // Copy constructor call
    y.SetStr( "asd" ); // will create string from "asd"
                       // and pass this to SetStr
    CallOnConst( y );
}

int main()
{
    TestNamespaces();
    TestString();
    TestMan();
    TestStrNum();
    TestConst();

    return 0;
}
