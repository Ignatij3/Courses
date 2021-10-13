#pragma once

class SomeData
{
    public:
        SomeData( int x );
        void Print();
    protected:
    private:
        int x_;
};

SomeData CreateSomeData();
