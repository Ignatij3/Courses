#include "Child.h"
#include <iostream>

using namespace std;

Child::Child() :
    Parent()
{
}

void Child::M1()
{
    cout << "Child::M1" << endl;
}

void Child::M2()
{
    cout << "Child::M2" << endl;
}
