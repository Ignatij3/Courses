#include <iostream>
#include <vector>
#include <algorithm>
#include "string.h"

using namespace std;

template< typename T >
T GetMin( const T &x1, const T &x2 )
{
    return ( x1 < x2 ) ? x1 : x2;
}

template< typename T >
void Swap( T &x1, T &x2 )
{
    T tmp = x1;
    x1 = x2;
    x2 = tmp;
}

template< typename T1, typename T2 >
int Compare( const T1 &x1, const T2 &x2 )
{
    if ( x1 < x2 )
    {
        return -1;
    }
    else if ( x1 > x2 )
    {
        return 1;
    }
    else
    {
        return 0;
    }
}

void TestFuncTemplate()
{
    cout << GetMin( 20, 10 ) << endl;
    double d1 = 10.0;
    double d2 = 9.5;
    cout << GetMin( d1, d2 ) << endl;
    cout << GetMin<double>( 10, 9.5 ) << endl;
    string s1 = "ABC";
    string s2 = "ABA";
    cout << GetMin( s1, s2 ) << endl;
    Swap( d1, d2 );
    cout << d1 << " " << d2 << endl;
    Swap( s1, s2 );
    cout << s1 << " " << s2 << endl;
    cout << Compare( 1, 2 ) << endl;
    cout << Compare( 1, 2.0 ) << endl;
    cout << Compare( s1, s2 ) << endl;
    // Will try to compile but fail cause
    // operator < in exprssion s1 < 2.0 is not defined
    // cout << Compare( s1, 2.0 ) << endl;
}

template< typename T >
T Avg( const T &x1, const T &x2 )
{
    return ( x1 + x2 ) / 2;
}

struct Point
{
    double x_;
    double y_;
    Point( double x, double y ) :
        x_( x ),
        y_( y )
    {
    }
    Point operator+( const Point &other ) const
    {
        return Point( x_ + other.x_, y_ + other.y_ );
    }
    Point operator/( double d ) const
    {
        return Point( x_ / d, y_ / d );
    }
};
ostream & operator <<( ostream &os, const Point &p )
{
    os << p.x_ << " " << p.y_ << endl;
    return os;
}

string operator/ ( const string &s1, int x )
{
    return s1.substr( 0, s1.length() / x );
}

void TestFuncTemplateWithClasses()
{
    cout << Avg( 2, 5 ) << endl;
    cout << Avg( 2.0, 5.0 ) << endl;
    cout << Avg( Point( 1, 2 ), Point( 3, 4 ) ) << endl;
    cout << Avg( string( "AB" ), string( "123456") ) << endl;
}

template< typename Dimension = double >
struct Point3D
{
    Dimension x_;
    Dimension y_;
    Dimension z_;
    Point3D( Dimension x, Dimension y, Dimension z ) :
        x_( x ),
        y_( y ),
        z_( z )
    {
    }
    template< typename T >
    Point3D operator+( const T &other )
    {
        return Point3D( x_ + other.x_, y_ + other.y_, z_ + other.z_ );
    }
};

template< typename DimensionType = double >
struct Point4D
{
    typedef DimensionType Dimension;
    Dimension x_;
    Dimension y_;
    Dimension z_;
    Dimension t_;
    // ...
};

template< typename Data, int Size = 128 >
struct Array
{
    Data d_[ Size ];
};

template< template< typename > class Point >
struct PointStorage
{
    vector< Point< double > > data_;

    void Add( const Point< double > &p );
};

template< template< typename > class Point >
void PointStorage<Point>::Add( const Point< double > &p )
{
    data_.push_back( p );
};

void TestClassTemplate()
{
    Point3D<> p1( 1, 2, 3);
    Point3D< float > p2( 4, 5, 6);
    cout << sizeof( p1 ) << endl;
    cout << sizeof( p2 ) << endl;

    Point3D<> p3 = p1 + p2;

    Array< char > a1;
    Array< char, 1024 > a2;
    cout << sizeof( a1 ) << endl;
    cout << sizeof( a2 ) << endl;

    PointStorage< Point3D > s1;
    PointStorage< Point4D > s2;
    cout << sizeof( s1.data_[0] ) << endl;
    cout << sizeof( s2.data_[0] ) << endl;

    s1.Add( Point3D<double>( 0, 1, 2 ) );

    Point4D<> p4;
    Point4D<>::Dimension d1 = p4.x_;
}

template<>
void Swap< int >( int &x1, int &x2 )
{
    // Will not use temporary variable
    x1 += x2;
    x2 = x1 - x2;
    x1 -= x2;
}
template<>
void Swap< vector< int > >( vector< int > &v1, vector< int > &v2 )
{
    v1.swap( v2 );
}

template<>
struct Point4D< char >
{
    // store 4 bytes in 1 integer
    int dim;
    // ...
};


void TestSpecTemplate()
{
    int i1 = 1;
    int i2 = 2;
    Swap( i1, i2 );
    cout << "Swap " << i1 << " " << i2  << endl;
    vector< int > v1(5, 10);
    vector< int > v2(3, 20);
    Swap( v1, v2 );
    Point4D< char > p;
    cout << sizeof( p ) << endl;
}

struct A
{
     int F1( double d )
     {
         cout << "F1" << d << endl;
         return d;
     }
     int F2( double d )
     {
         cout << "F2" << d << endl;
         return d;
     }
     int F3( double d, const string &s )
     {
         cout << "F3" << d << s << endl;
         return d;
     }
};
typedef int (A::* FCallback)( double d );

#include <functional>

typedef std::function< int( double d ) > StdCallback;
void Caller( StdCallback &c )
{
    c( 50 );
}

void TestCallback()
{
    // pointer to member functions
    FCallback c1 = &A::F1;
    FCallback c2 = &A::F2;
    A a;
    int x1 = (a.*c1)( 10 );
    int x2 = (a.*c2)( 20 );

    // std::function
    function< int( A&, double ) > c3 = &A::F1;
    int x3 = c3( a, 30 );

    // std::bind std::placeholders
    function< int( double d ) > c4 = bind( &A::F1, &a, placeholders::_1 );
    function< int( double d ) > c5 = bind( &A::F2, &a, placeholders::_1 );
    c4(40);
    Caller( c4 );
    Caller( c5 );

    function< int( A&, double d, const string &s ) > c6 = &A::F3;
    c6( a, 60, string( "aaa" ) );

    function< int( double d, const string &s ) > c7 = bind( &A::F3, &a, placeholders::_1, placeholders::_2 );
    c7( 70, string( "aaa" ) );

    string s2 = "bbb";
    function< int( double d ) > c8 = bind( &A::F3, &a, placeholders::_1, s2 );
    c8( 80 );
}

//template <class T> struct equal_to : binary_function <T,T,bool> {
//  bool operator() (const T& x, const T& y) const {return x==y;}
//};

template <class T>
struct similar_to : binary_function< T, T, bool >  {
  bool operator() (const T& x, const T& y) const
  {
      return abs(x-y) <= 10;
  }
};

void TestBinders()
{
    int a1[] = {10, 20, 32, 20, 20};
    int a2[] = {10, 20, 40, 80, 160};
    std::pair< int *, int * > r =
        mismatch(a1, a1 + 5, a2, equal_to<int>() );
    cout << "mismatch " << *r.first <<  " " << *r.second << endl;
    r = mismatch(a1, a1 + 5, a2, similar_to<int>() );
    cout << "mismatch " << *r.first <<  " " << *r.second << endl;

    int c1 = count_if( a1, a1 + 5, bind1st( equal_to<int>(), 20 ) );
    int c2 = count_if( a1, a1 + 5, bind( equal_to<int>(), 20, placeholders::_1 ) );
    int c3 = count_if( a1, a1 + 5, bind( similar_to<int>(), 20, placeholders::_1 ) );
    cout << c1 << " " << c2 << " " << c3 << " elements is 20" << endl;

}

int main()
{
    TestFuncTemplate();
    TestFuncTemplateWithClasses();
    TestClassTemplate();
    TestSpecTemplate();
    TestCallback();
    TestBinders();

    return 0;
}
