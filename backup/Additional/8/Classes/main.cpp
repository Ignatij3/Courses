#include <iostream>
#include <cmath>
#include <cstdio>
#include <vector>
#include <list>
using namespace std;

#include <map>
void TestMap()
{
    typedef map<string, int> MyMap;
    typedef map<string, int>::iterator MyMapIt;
    typedef map<string, int>::const_iterator MyMapCIt;
    typedef pair<string, int> MyMapEl;
    MyMap m1;

    m1["key1"] = 100;
    m1["key2"] = 200;
    for( MyMapIt it = m1.begin(); it != m1.end(); ++it)
    {
        cout << it->first << " = " << it->second << endl;
    }
    cout << m1["key1"] << endl;
    // BAD compare
    if ( m1["key3"] == 300 )
    {
        cout << "key3 yes 300" << endl;
    }
    if ( m1["key3"] == 0 )
    {
        cout << "key3 yes 0" << endl;
    }
    if (  m1.find( "key4" ) == m1.end() )
    {
        cout << "key 4 not found" << endl;
    }
    MyMapCIt itKey1 = m1.find( "key1" );
    if ( itKey1 != m1.end() )
    {
        cout << "key 1 found" << endl;
        if ( itKey1->second == 100 )
        {
            cout << "and equal 100" << endl;
        }
    }
    m1["key2"] = 201;
    MyMapIt itKey2 = m1.find( "key2" );
    if ( itKey2 != m1.end() )
    {
        itKey2->second = 202;
    }
    cout << m1["key2"] << endl;
}
/*
empty Test whether container is empty
size Return container size
insert Insert elements
erase Erase elements
swap Swap content
clear Clear content
find Get iterator to element
count Count elements with a specific key
*/

#include <set>
void TestSet()
{
    typedef set<string> MySet;
    typedef MySet::iterator MySetIt;
    typedef MySet::const_iterator MySetCIt;
    //typedef pair<string, int> MyMapEl;

   MySet data;
   data.insert("key1");
   data.insert("key2");
   data.insert("key1");
   for( MySetCIt it = data.begin(); it != data.end(); ++it)
   {
       cout << *it << endl;
   }
   data.erase("key2");
   cout << data.size() << endl;
   data.insert("key3");
   if ( data.count("key3") != 0 )
   {
       cout << "found" << endl;
   }
   MySetIt it = data.find("key3");
   if ( it != data.end() )
   {
       data.erase( it );
       cout << data.size() << endl;
   }
}
/*
empty Test whether container is empty
size Return container size
insert Insert element
erase Erase elements
swap Swap content
clear Clear content
find Get iterator to element
count Count elements with a specific value
*/

#include <bitset>
void TestBitSet()
{
    enum
    {
        OptCanRead,
        OptCanWrite,
        OptCanDelete,
        OptCanCreate,
        OptSize
    };

    bitset<OptSize> user1;
    user1.set( OptCanRead );
    user1.set( OptCanWrite );
    bitset<OptSize> user2;
    user1.set( OptCanRead );
    if ( user1[ OptCanRead ] )
    {
        cout << "you can read" << endl;
    }
    if ( !user2[ OptCanWrite ] )
    {
        cout << "you can't write" << endl;
    }
}

#include <algorithm>
#include <numeric>
#include <functional>
bool ifMod3Func( int i )
{
    return ( ( i % 3 ) == 0 );
}

class Data
{
friend class Printer;
public:
    struct Printer
    {
        void operator()( const Data &d )
        {
            cout << d.name_ << endl;
        }
    };
    static void print( const Data &d )
    {
        cout << d.name_ << endl;
    }
    Data( const string &name ) :
        name_( name )
    {
    }

private:
    string name_;
};

void TestAlgorithm()
{
    typedef vector<int> MyVector;
    typedef MyVector::iterator MyVectorIt;
    int cInts[] = { 10, 20, 30 ,40 };
    int cIntsSize = sizeof( cInts ) / sizeof( cInts[0] );
    // Works with C-arrays
    int *p = find( cInts, cInts + cIntsSize, 30 );
    cout << "found " << *p << endl;

    MyVector vInts( cInts, cInts + cIntsSize );
    MyVectorIt it = find( vInts.begin(), vInts.end(), 30 );
    if ( it != vInts.end() )
    {
        cout << "found " << *it << endl;
    }

    it = find_if( vInts.begin(), vInts.end(), ifMod3Func );
    cout << "first mod 3 is " << *it << endl;

    struct IfMod3
    {
        bool operator()( int i )
        {
            return ( ( i % 3 ) == 0 );
        }
    };
    it = find_if( vInts.begin(), vInts.end(), IfMod3() );
    cout << "first mod 3 is " << *it << endl;

    MyVector v20and30(2);
    v20and30[0] = 20;
    v20and30[1] = 30;
    it = search( vInts.begin(), vInts.end(), v20and30.begin(), v20and30.end() );
    cout << "20 and 30 found at " << distance( vInts.begin(), it ) << endl;

    vector< Data > data;
    data.push_back( Data( "asd" ) );
    data.push_back( Data( "qwe" ) );
    for_each( data.begin(), data.end(), Data::Printer() );
    for_each( data.begin(), data.end(), Data::print );

    struct FibNumberGen
    {
        int a1_;
        int a2_;
        FibNumberGen() :
            a1_(0),
            a2_(0)
        {
        }
        int operator() ()
        {
            if ( a1_ == 0 )
            {
                a1_ = 1;
                return 1;
            }
            if ( a2_ == 0 )
            {
                a2_ = 1;
                return 1;
            }
            int x = a1_ + a2_;
            a1_ = a2_;
            a2_ = x;
            return x;
        }
    };
    MyVector v2( 10 );
    generate( v2.begin(), v2.end(), FibNumberGen());
    for( MyVectorIt it = v2.begin(); it != v2.end(); ++it )
    {
        cout << *it << endl;
    }

    set< int > s1;
    set< int > s2;
    set< int > s3;
    s1.insert(10); s1.insert(20); s1.insert(30);
    s2.insert(13); s2.insert(10); s2.insert(20);
    set_intersection( s1.begin(), s1.end(), s2.begin(), s2.end(),
                      inserter( s3, s3.begin() ) );
    for( set< int >::const_iterator it = s3.begin(); it != s3.end(); ++it )
    {
        cout << *it << endl;
    }

    cout << accumulate( s3.begin(), s3.end(), 0, plus< int >() ) << endl;

}

int add_to_totals( int total, const map< std::string, int >::value_type& data )
{
    return total + data.second;
}

int main()
{
    map<string, int > m;
    m[ "a" ] = 1;
    m[ "b" ] = 2;
    const int total = std::accumulate( m.begin(), m.end(), 0, add_to_totals );
    cout << total << endl;
    TestMap();
    TestSet();
    TestBitSet();
    TestAlgorithm();

    return 0;
}
