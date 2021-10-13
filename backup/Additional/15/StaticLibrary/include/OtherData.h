#pragma once

class OtherData
{
    public:
        OtherData( int x );
        void Print();
    protected:
    private:
        int x_;
};

OtherData CreateOtherData();
