#include <iostream>
#include <cmath>
#include <cstdio>
using namespace std;

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

#include <vector>
void TestVector()
{
    int aInts[1000];
    vector< int > vInts(1000);
    cout << sizeof( aInts ) << endl;
    cout << sizeof( vInts ) << endl; // 12 bytes
    for( int i = 0; i < 10; ++i )
    {
        vInts[i] = i * i;
        cout << vInts[i] << endl;
    }
    vector <int> v; // 0 size
    cout << v.size() << " " << v.capacity() << endl;
    for( int i = 0; i < 10; ++i )
    {
        v.push_back( i );
        cout << v.size() << " " << v.capacity() << endl;
    }
    cout << v.front() << " " << v.back() << endl;
    /*  0 0
        1 1
        2 2
        3 4
        4 4
        5 8
        6 8
        7 8
        8 8
        9 16
        10 16
        0 9 */
    vInts = v;
    cout << vInts.back() << endl; // 9
}
/*
size Return size
empty Test whether vector is empty
resize Change size
reserve Request a change in capacity
shrink_to_fit Shrink to fit(Cx11 only)
front Access first element
back Access last element
push_back Add element at the end
pop_back Delete last element
insert Insert elements
erase Erase elements
swap Swap content(sizes may differ)
clear Clear content
*/

#include <list>
#include <iterator>
void TestList()
{
    list< int > ints;
    for( int i = 0; i < 5; ++i )
    {
        ints.push_front( i );
    }
    // for 2
    for( list<int>::const_iterator it = ints.begin();
         it != ints.end(); ++it )
    {
        cout << *it << endl;
    }
    // for 3
    for( list<int>::iterator it = ints.begin();
         it != ints.end(); ++it )
    {
        *it =  *it * 2;
        cout << *it << endl;
    }
    // for 4, bad iterators can be iterated back
    // prints <uninitialized memory, e.g 5> 0 2 4 6
    for( list<int>::const_iterator it = ints.end();
         it != ints.begin(); --it )
    {
        cout << *it << endl;
    }
    // for 5, good, use reverse iterator
    for( list<int>::const_reverse_iterator rit = ints.rbegin();
         rit != ints.rend(); ++rit )
    {
        cout << *rit << endl;
    }
    // for 6
    list< int >::iterator it3 = ints.begin();
    advance( it3, 2 );
    list< int > l2( it3, ints.end());
    for( list<int>::const_iterator it = l2.begin();
         it != l2.end(); ++it )
    {
        cout << *it << endl;
    }
    // for 7
    vector< int > v2(ints.rbegin(), ints.rend());
    for( vector< int >::const_iterator it = v2.begin();
         it != v2.end(); ++it )
    {
        cout << *it << endl;
    }
    // for 8-9
    vector<int> v3( 4, 100 );// 4 ints with 100 value
    v3.insert( v3.begin() + 2, 200 );
    for( vector< int >::const_iterator it = v3.begin();
         it != v3.end(); ++it )
    {
        cout << *it << endl;
    }
    list<int> l3(2, 300);
    l3.insert( l3.begin(), v3.begin(), v3.end());
    for( list< int >::const_iterator it = l3.begin();
         it != l3.end(); ++it )
    {
        cout << *it << endl;
    }
    // for 10-11
    for( list< int >::iterator it = l3.begin();
         it != l3.end(); ++it )
    {
        l3.insert( it, 13 );
    }
    for( list< int >::const_iterator it = l3.begin();
         it != l3.end(); ++it )
    {
        cout << *it << endl;
    }
}
/*
empty Test whether container is empty
size Return size
assign Assign new content to container
push_front Insert element at beginning
pop_front Delete first element
push_back Add element at the end
pop_back Delete last element
insert Insert elements
erase Erase elements
swap Swap content
resize Change size
clear Clear content
remove Remove elements with specific value
unique Remove duplicate values
sort Sort elements in container
merge Merge sorted lists
reverse Reverse the order of elements */

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
    for( MyMapCIt it = m1.begin(); it != m1.end(); ++it)
    {
        MyMapEl el = *it;
        cout << el.first << " = " << el.second << endl;
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

class A
{
public:
    virtual ~A() {};
    virtual void PrintMe() = 0;
};
class B : public A
{
    virtual void PrintMe()
    {
        cout << "It is B at " << (unsigned long int)this << endl;
    }
};
class C : public A
{
    virtual void PrintMe()
    {
        cout << "It is C at " << (unsigned long int)this << endl;
    }
};

void TestPolymorphism()
{
    typedef vector< A * > Objs;
    Objs objs;
    objs.push_back( new B() );
    objs.push_back( new C() );
    for( int i = 0; i < objs.size(); ++i )
    {
        objs[i]->PrintMe();
    }
    for( int i = 0; i < objs.size(); ++i )
    {
        delete objs[i];
    }
    objs.clear();
    cout << objs.size();
}

void TestPolymorphismIterators()
{
    typedef list< A * > Objs;
    typedef Objs::iterator ObjsIt;

    Objs objs;
    objs.push_back( new B() );
    objs.push_back( new C() );
    for( ObjsIt it = objs.begin(); it != objs.end(); ++it )
    {
        A *a = *it;
        a->PrintMe();

        A &b = **it;
        b.PrintMe();

        delete a;
    }
}


int main()
{
    /*
    TestConst();
    TestVector();
    TestList();
    TestMap();
    TestPolymorphism();
    */
    TestPolymorphismIterators();

    return 0;
}
