#include <iostream>
#include <cmath>

using namespace std;

class Book
{
public:
    Book()
    {
        ++count_;
    }
    ~Book()
    {
        --count_;
    }
    static int GetCount();
    //static int GetCount()
    //{
    //   return count_;
    //}
private:
    static int count_;

};
int Book::count_ = 0;
int Book::GetCount()
{
    return count_;
}

void TestStatic()
{
    cout << Book::GetCount() << endl;
    Book b;
    cout << Book::GetCount() << endl;
    {
        Book bb[10];
        cout << Book::GetCount() << endl;
    }
    cout << Book::GetCount() << endl;
}

class Math
{
public:
    static double Sqr( double x )
    {
        return x * x;
    }
    static double C( double a, double b )
    {
        return sqrt( Sqr( a ) + Sqr( b ) );
    }
};

void TestStaticClass()
{
    cout << Math::C( 3, 4 ) << endl;

}

class Preferences
{
public:
    static Preferences &Instance()
    {
        // Guaranteed to be lazy initialized
        static Preferences instance;
        return instance;
    }

    int GetMaxUsers()
    {
        return maxUsers_;
    }

private:
    Preferences() :
        maxUsers_( 10 )
    {};
    int maxUsers_;
};

void TestSingleton()
{
    cout << Preferences::Instance().GetMaxUsers() << endl;
}


class Figure
{
public:
    virtual ~Figure(){};
    virtual void Draw() = 0;
};

class Square : public Figure
{
public:
    Square( double side, double color ) :
        side_( side ),
        color_( color )
    {
    }
    virtual void Draw()
    {
        cout << "Square draw" << endl;
    }
private:
    double side_;
    double color_;
};

class Circle : public Figure
{
public:
    Circle( double r ) :
        r_( r )
    {
    }
    virtual void Draw()
    {
        cout << "Circle draw" << endl;
    }
private:
    double r_;
};

void Draw( Figure &f )
{
    f.Draw();
}

class FigureFactory
{
public:
    static Figure *CreateCircle( double r )
    {
        return new Circle( r );
    }
    static Figure *CreateEarth()
    {
        return new Circle( 63e9  );
    }
    static Figure *CreateRedSquare( double side )
    {
        return new Square( side, 0xFF0000 );
    }
    static Figure *CreateSmallGreenSquare()
    {
        return new Square( 1e-5, 0x00FF00 );
    }
};

void TestFactory()
{
    Figure *f[5];
    f[0] = new Circle( 10 );
    f[1] = FigureFactory::CreateCircle( 10 );
    f[2] = FigureFactory::CreateEarth();
    f[3] = FigureFactory::CreateRedSquare( 10 );
    f[4] = FigureFactory::CreateSmallGreenSquare();
    for( int i = 0; i < 5; ++i )
    {
        f[i]->Draw();
        delete f[i];
        f[i] = 0;
    }
}

class A
{
    friend class B;
    friend void Foo( A &a );
private:
    int x_;
    void ChangeX()
    {
        x_ = 10;
    }
};
class B
{
public:
    static void ChangeA( A &a )
    {
        a.x_ = 12;
        cout << a.x_ << endl;
    }
};

void Foo( A &a )
{
    a.ChangeX();
    cout << a.x_ << endl;
}

void TestFriends()
{
    A a;
    B::ChangeA( a );
    Foo( a );
}

int main()
{
    TestStatic();
    TestStaticClass();
    TestSingleton();
    TestFactory();
    TestFriends();
    return 0;
}
