#include "OtherData.h"
#include <iostream>

using namespace std;

OtherData::OtherData( int x ) :
    x_( x )
{
}

void OtherData::Print()
{
    cout << "OtherData::x = " << x_ << endl;
}

OtherData CreateOtherData()
{
    return OtherData( 100 );
}

