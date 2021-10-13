#include <iostream>
#include <cmath>
#include <cstdio>
#include <vector>
#include <list>
using namespace std;

struct Person
{
    string name_;
    string surname_;
};
ostream & operator << ( ostream &os, const Person &p)
{
    os << p.name_ << " " << p.surname_;
    return os;
}

istream & operator >> ( istream &is, Person &p)
{
    is >> p.name_;
    is >> p.surname_;
    // when error
    // is.setstate(ios::failbit);
    return is;
}

void TestStdStream()
{
    cout << "stdout" << endl;
    cerr << "stderr" << endl;
    string s, s2;
    cin >> s; // stdin

    cout << "word1 " << "word2" << endl;
    cin >> s >> s2;

    Person p;
    p.name_ = "John";
    p.surname_ = "Smith";
    cout << p << endl;
    Person p2;
    cin >> p2;
    cout << p2 << endl;

}

#include <fstream>
void TestFileStreamIn()
{
    // "d1.txt: 12345"
    fstream f1 ("d1.txt", fstream::in);
    if (!f1)
    {
        cout << "can't open d1.txt" << endl;
        return;
    }
    char c;
    while ( f1.get( c ) )
    {
        cout << "Get : " << c << endl;
    }
    f1.close();

    fstream f2 ("d_lines.txt", fstream::in);
    string line;
    while ( getline( f2, line ) )
    {
        cout << "Get line: " << line << endl;
    }
    f2.close();

    fstream f3 ("d2.txt", fstream::in);
    int x;
    while ( f3 >> x )
    {
        cout << "Get: " << x << endl;
    }
    f3.close();
}

#include <iomanip>
void TestFileStreamOut()
{
    fstream f1 ("d_out.txt", fstream::out);
    f1.put('A');
    f1 << 'B' << endl;
    string s =  "x=";
    int x  = 123;
    f1 << s << x << endl;
    double d  = 1.2345;
    f1 << setprecision(2) << d << endl;
    f1.close();

    fstream f2 ("d_out.txt", fstream::out | fstream::app );
    f2 << "More data" << endl;
    f2.close();
}

void TestFileBinary()
{
    vector<int> data(10);
    for(unsigned int i=0; i < data.size(); ++i )
    {
        data[i] = i*i;
    }
    fstream f1("d.bin", fstream::out | fstream::binary);
    f1.write( (char *)&data[0], data.size() * sizeof(int) );
    f1.close();

    fstream f2("d.bin", fstream::in | fstream::binary);
    int x;
    while ( f2.read( (char *)&x, sizeof(x) ) )
    {
        cout << x << endl;
    }
    f2.close();


    fstream f3( "d.bin", fstream::in | fstream::binary | fstream::ate);
    int fileSize = f3.tellg();
    int intsCount = fileSize / sizeof(int);
    vector<int> v( intsCount );
    f3.seekg( 0, fstream::beg );
    f3.read( (char *)&v[0], fileSize );
    f3.close();
    for(unsigned int i=0; i < v.size(); ++i )
    {
        cout << v[i] << endl;
    }
}

#include <sstream>
void TestStringStream()
{
    stringstream ss;
    ss << 110 << " " << "120" << " " << true;
    string s;
    int x ;
    bool b;
    ss >> s >> x >> b;

    cout << s << " " << x << " " << b << endl;
    cout << ss.str() << endl;
}

int main()
{
    //TestStdStream();
    TestFileStreamIn();
    TestFileStreamOut();
    TestFileBinary();

    TestStringStream();
    return 0;
}
