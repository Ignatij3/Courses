#include "SomeData.h"
#include <iostream>

using namespace std;

SomeData::SomeData( int x ) :
    x_( x )
{
}

void SomeData::Print()
{
    cout << "SomeData::x = " << x_ << endl;
}

SomeData CreateSomeData()
{
    return SomeData( 100 );
}

