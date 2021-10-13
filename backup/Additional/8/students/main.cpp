#include <iostream>
#include <string>
#include <algorithm>
#include <map>

using namespace std;

int add_to_totals( int total, const map< std::string, int >::value_type &data )
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

    return 0;
}
