#include <iostream>
#include "AllegroBase.hpp"
#include <windows.h>
#include <cstdlib>
#include <cstdio>

using namespace std;

const int FPS = 60;
const int SCREEN_W = 640;
const int SCREEN_H = 480;

class Figure
{
protected:
    double x_;
    double y_;
    double dx_;
    double dy_;

public:
    Figure()
    {
        Reset();
    }

    virtual ~Figure(){};

    void Reset()
    {
        x_ = rand() % SCREEN_W;
        y_ = rand() % SCREEN_H;
        dx_ = 10.0 - rand() % 21;
        dy_ = 10.0 - rand() % 21;
    }

    virtual void Draw() = 0;

    virtual void Move()
    {
        x_ += dx_;
        y_ += dy_;
        if ( ( x_ < 1.0 ) ||
             ( x_ > SCREEN_W ) ||
             ( y_ < 1.0 ) ||
             ( y_ > SCREEN_H ) )
        {
            Reset();
        }
    };

};
typedef Figure * PFigure;

class Square : public Figure
{
protected:
    double a_;
public:
    Square( double a ) :
        Figure(),
        a_( a )
    {
    }
    virtual void Draw()
    {
        double half = a_ / 2;
        al_draw_filled_rectangle( x_ - half, y_ - half,
            x_ + half, y_ + half, al_map_rgb( 255, 0, 0 ) );
    }
};

class Circle : public Figure
{
protected:
    double r_;
    unsigned char color_;
public:
    Circle( double r ) :
        Figure(),
        r_( r ),
        color_( rand() % 256 )
    {
    }
    virtual void Draw()
    {
        ++color_;
        al_draw_filled_circle( x_, y_, r_, al_map_rgb( 0, color_, 0 ) );
    }
};

const int MAX = 100;
class ScreenSaver
{
private:
    PFigure figures[MAX];
    int size_;

public:
    static ScreenSaver &Instance()
    {
        static ScreenSaver instance;
        return instance;
    }

    void Draw()
    {
        al_clear_to_color( al_map_rgb( 0, 0, 0 ) );
        for( int i = 0; i < size_; ++i )
        {
            figures[i]->Draw();
        }
    }

    void Next()
    {
        for( int i = 0; i < size_; ++i )
        {
            figures[i]->Move();
        }
    }

    void Add( Figure *f )
    {
        if ( size_ >= MAX )
        {
            return;
        }
        figures[ size_ ] = f;
        ++size_;
    }

    void Reset()
    {
        for( int i = 0; i < size_; ++i )
        {
            figures[i]->Reset();
        }
    }

private:
    ScreenSaver() :
        size_( 0 )
    {
        // Set to null all pointers
        memset( figures, 0, sizeof( figures ) );
    }

    ~ScreenSaver()
    {
        for( int i = 0; i < size_; ++i )
        {
            delete figures[i];
            figures[i] = 0;
        }
    }

};

class ControlledSquare : public Square
{
public:
    ControlledSquare( double side ) :
        Square( side )
    {
    }

    void MoveBy( double dx, double dy )
    {
        dx_ = dx;
        dy_ = dy;
        Move();
    }
};

class FigureFactory
{
public:
    enum Type
    {
        RandomCircle,
        RandomSquare
    };

    static Figure * Create( Type type )
    {
        switch( type )
        {
            case RandomCircle:
                return new Circle( 10.0 + rand() % 30 );
            case RandomSquare:
                return new Square( 10.0 + rand() % 30 );
        }
    }
};

class AllegroApp : public AllegroBase
{
private:
    ControlledSquare humanSquare_;

public:
    AllegroApp() :
        AllegroBase(),
        humanSquare_( 30 )
    {
        for( int i = 0; i < 10; ++i )
        {
            if ( ( i % 2 ) == 0 )
            {
                ScreenSaver::Instance().Add(
                    FigureFactory::Create( FigureFactory::RandomSquare ) );
            }
            else
            {
                ScreenSaver::Instance().Add(
                    FigureFactory::Create( FigureFactory::RandomCircle ) );
            }
        }
    }

    virtual void Fps()
    {
        ScreenSaver::Instance().Next();

        double dx = 0, dy = 0;
        if ( IsPressed( ALLEGRO_KEY_UP ) )
        {
            dy = -1;
        }
        if ( IsPressed( ALLEGRO_KEY_DOWN ) )
        {
            dy = +1;
        }
        if ( IsPressed( ALLEGRO_KEY_LEFT ) )
        {
            dx = -1;
        }
        if ( IsPressed( ALLEGRO_KEY_RIGHT ) )
        {
            dx = +1;
        }
        if ( IsPressed( ALLEGRO_KEY_LSHIFT ) )
        {
            dy *= 10;
            dx *= 10;
        }
        humanSquare_.MoveBy( dx, dy );
    }

    virtual void Draw()
    {
        ScreenSaver::Instance().Draw();
        humanSquare_.Draw();
    }

    virtual void OnKeyDown( const ALLEGRO_KEYBOARD_EVENT &keyboard )
    {
        if ( keyboard.keycode == ALLEGRO_KEY_SPACE )
        {
            ScreenSaver::Instance().Reset();
        }
        else if ( keyboard.keycode == ALLEGRO_KEY_ESCAPE )
        {
            Exit();
        }
    }
};

int main(int argc, char **argv)
{
    srand( time(0) );

    AllegroApp app;
    if ( !app.Init( SCREEN_W, SCREEN_H, FPS ) )
    {
        return 1;
    }
    app.Run();

    // cin.get();
    return 0;
}
