#include <iostream>
#include "AllegroBase.hpp"
#include <windows.h>
#include <cstdlib>
#include <cstdio>
#include <fstream>
#include <vector>

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
    virtual string GetType() const = 0;
    virtual void Save( ostream &os ) const = 0;
    virtual void Load( istream &is ) = 0;

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

ostream & operator << ( ostream &os, const Figure &f )
{
    f.Save( os );
    return os;
}

istream & operator >> ( istream &is, Figure &f )
{
    f.Load( is );
    return is;
}

class Square : public Figure
{
protected:
    double a_;
public:
    Square() :
        Figure(),
        a_( 0 )
    {
    }
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
    virtual string GetType() const
    {
        return "square";
    }
    virtual void Save( ostream &os ) const
    {
        os << x_ << " " << y_ << " " << dx_ << " " << dy_ << " "
           << a_;
    }
    virtual void Load( istream &is )
    {
        is >> x_ >> y_ >> dx_ >> dy_ >> a_;
    }
};

class Circle : public Figure
{
protected:
    double r_;
    unsigned char color_;
public:
    Circle() :
        Figure(),
        r_( 0 ),
        color_( 0 )
    {
    }
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
    virtual string GetType() const
    {
        return "circle";
    }
    virtual void Save( ostream &os ) const
    {
        os << x_ << " " << y_ << " " << dx_ << " " << dy_ << " "
           << r_ << " " << color_;
    }
    virtual void Load( istream &is )
    {
        is >> x_ >> y_ >> dx_ >> dy_ >> r_ >> color_;
    }
};

class ScreenSaver
{
private:
    vector<PFigure> figures;
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
        for( int i = 0; i < figures.size(); ++i )
        {
            figures[i]->Draw();
        }
    }

    void Next()
    {
        for( int i = 0; i < figures.size(); ++i )
        {
            figures[i]->Move();
        }
    }

    void Add( Figure *f )
    {
        figures.push_back( f );
    }

    void Reset()
    {
        for( int i = 0; i < figures.size(); ++i )
        {
            figures[i]->Reset();
        }
    }

    void Save( ostream &os )
    {
        for( int i = 0; i < figures.size(); ++i )
        {
            os << figures[i]->GetType() << " "
               << *(figures[i]) << endl;
        }
    }
    void Load( istream &is )
    {
        Clear();

        string type;
        while( is >> type )
        {
            Figure *f;
            if ( type == "square" )
            {
                f = new Square();
            }
            else
            {
                f = new Circle();
            }
            is >> *f;
            Add( f );
        }
    }

private:
    ScreenSaver() :
        figures()
    {
    }

    ~ScreenSaver()
    {
        Clear();
    }
    void Clear()
    {
        for( int i = 0; i < figures.size(); ++i )
        {
            delete figures[i];
        }
        figures.clear();
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
        else if ( keyboard.keycode == ALLEGRO_KEY_S )
        {
            fstream fs( "save.txt", fstream::out );
            ScreenSaver::Instance().Save( fs );
        }
        else if ( keyboard.keycode == ALLEGRO_KEY_L )
        {
            fstream fs( "save.txt", fstream::in );
            ScreenSaver::Instance().Load( fs );
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
