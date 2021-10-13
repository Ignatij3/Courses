#include "AllegroBase.hpp"

#include <cmath>
#include <cstdio>
#include <cstdlib>
#include <iostream>
#include <windows.h>

using namespace std;

const int FPS      = 60;
const int SCREEN_W = 640;
const int SCREEN_H = 480;

class Figure {
  protected:
    double x_;
    double y_;
    double dx_;
    double dy_;

  public:
    Figure() :
        x_(0),
        y_(0),
        dx_(0),
        dy_(0)
    {
    }

    virtual ~Figure() {};

    virtual void Reset()
    {
        x_  = rand() % SCREEN_W;
        y_  = rand() % SCREEN_H;
        dx_ = 10.0 - rand() % 21;
        dy_ = 10.0 - rand() % 21;
    }

    virtual void Draw() = 0;

    virtual void Move()
    {
        x_ += dx_;
        y_ += dy_;
        if ((x_ < 1.0) || (x_ > SCREEN_W) || (y_ < 1.0) || (y_ > SCREEN_H))
        {
            Reset();
        }
    };
};
typedef Figure* PFigure;

class Square : public Figure {
  protected:
    double a_;

  public:
    Square(double a) :
        Figure(),
        a_(a)
    {
    }
    virtual void Draw()
    {
        double half = a_ / 2;
        al_draw_filled_rectangle(x_ - half, y_ - half,
            x_ + half, y_ + half, al_map_rgb(255, 0, 0));
    }
};

class Circle : public Figure {
  protected:
    double r_;
    unsigned char cR_, cG_, cB_;

  public:
    Circle(double r, unsigned char cR, unsigned char cG, unsigned char cB) :
        Figure(),
        r_(r),
        cR_(cR),
        cG_(cG),
        cB_(cB)
    {
    }
    virtual void Draw()
    {
        al_draw_filled_circle(x_, y_, r_, al_map_rgb(cR_, cG_, cB_));
    }
};

const int MAX = 100;
class ScreenSaver {
  private:
    PFigure figures[MAX];
    int size_;

  public:
    static ScreenSaver& Instance()
    {
        static ScreenSaver instance;
        return instance;
    }

    void Draw()
    {
        al_clear_to_color(al_map_rgb(0, 0, 0));
        for (int i = 0; i < size_; ++i)
        {
            figures[i]->Draw();
        }
    }

    void Next()
    {
        for (int i = 0; i < size_; ++i)
        {
            figures[i]->Move();
        }
    }

    void Add(Figure* f)
    {
        if (size_ >= MAX)
        {
            return;
        }
        figures[size_] = f;
        ++size_;
    }

    void Reset()
    {
        for (int i = 0; i < size_; ++i)
        {
            figures[i]->Reset();
        }
    }

  private:
    ScreenSaver() :
        size_(0)
    {
        // Set to null all pointers
        memset(figures, 0, sizeof(figures));
    }

    ~ScreenSaver()
    {
        for (int i = 0; i < size_; ++i)
        {
            delete figures[i];
            figures[i] = 0;
        }
    }
};

class Body : public Square {
    friend class Robot;

  public:
    Body(double side) :
        Square(side)
    {
        Reset();
    }

    void MoveBy(double dx, double dy)
    {
        dx_ = dx;
        dy_ = dy;
        Move();
    }
};

class Projectile : public Circle {
    friend class Gun;

  private:
    bool fired_;
    double angle_;
    double speed_;

  public:
    Projectile() :
        Circle(5, 100, 100, 100),
        fired_(false),
        angle_(0),
        speed_(0)
    {
    }

    virtual void Reset()
    {
        x_     = 0;
        y_     = 0;
        fired_ = false;
    }

    virtual void DrawAt(double stX, double stY)
    {
        if (not fired_)
        {
            x_ = stX;
            y_ = stY;
        }
        Draw();
    }

    void Process()
    {
        if (fired_)
        {
            // Not so good, angle can be changed when turning gun
            dx_ = speed_ * cos(angle_);
            dy_ = speed_ * sin(angle_);
            Move();
        }
    }
};

class Gun {
  private:
    double angle_;
    double r_;
    double speed_;
    Projectile projectile_;

  public:
    Gun(double r) :
        angle_(0),
        r_(r),
        projectile_(),
        speed_()
    {
    }
    void Draw(double stX, double stY)
    {
        double nX = endX(stX);
        double nY = endY(stY);
        al_draw_line(stX, stY, nX, nY,
            al_map_rgb(0, 0, 255), 3);
        projectile_.DrawAt(nX, nY);
    }
    double endX(double stX)
    {
        return stX + r_ * cos(angle_);
    }
    double endY(double stY)
    {
        return stY + r_ * sin(angle_);
    }
    void SetAngle(double angle)
    {
        angle_ = angle;
    }
    void Fire()
    {
        projectile_.fired_ = true;
        projectile_.angle_ = angle_;
        projectile_.speed_ = 10;
    }
    void Process()
    {
        projectile_.Process();
    }
};

struct Target {
    double x, y;
    Target(double x0, double y0) :
        x(x0), y(y0)
    {
    }
};

class Robot {
  private:
    Body body_;
    Gun gun_;
    Target target_;

  public:
    Robot(double size) :
        body_(size),
        gun_(size * 1.15),
        target_(gun_.endX(body_.x_), gun_.endY(body_.y_))
    {
    }

    void Process(double dx, double dy)
    {
        if (dx != 0 || dy != 0)
        {
            body_.MoveBy(dx, dy);
            TurnGunToTarget();
        }
        gun_.Process();
    }
    void Draw()
    {
        body_.Draw();
        gun_.Draw(body_.x_, body_.y_);
    }
    void SetTarget(double x, double y)
    {
        target_.x = x;
        target_.y = y;
    }
    void TurnGunToTarget()
    {
        double a = target_.x - body_.x_;
        double b = target_.y - body_.y_;
        gun_.SetAngle(atan2(b, a));
    }
    void Fire()
    {
        gun_.Fire();
    }
};

class FigureFactory {
  public:
    enum Type {
        RandomCircle,
        RandomSquare
    };

    static Figure* Create(Type type)
    {
        Figure* f;
        switch (type)
        {
        case RandomCircle:
            f = new Circle(10.0 + rand() % 30,
                0, 255, 0);
            break;
        case RandomSquare:
            f = new Square(10.0 + rand() % 30);
            break;
        }
        f->Reset();
        return f;
    }
};

class AllegroApp : public AllegroBase {
  private:
    Robot humanRobot_;

  public:
    AllegroApp() :
        AllegroBase(),
        humanRobot_(30)
    {
        for (int i = 0; i < 10; ++i)
        {
            if ((i % 2) == 0)
            {
                ScreenSaver::Instance().Add(
                    FigureFactory::Create(FigureFactory::RandomSquare));
            }
            else
            {
                ScreenSaver::Instance().Add(
                    FigureFactory::Create(FigureFactory::RandomCircle));
            }
        }
    }

    virtual void Fps()
    {
        ScreenSaver::Instance().Next();
        double dx = 0, dy = 0;
        if (IsPressed(ALLEGRO_KEY_W))
        {
            dy = -1;
        }
        if (IsPressed(ALLEGRO_KEY_S))
        {
            dy = +1;
        }
        if (IsPressed(ALLEGRO_KEY_A))
        {
            dx = -1;
        }
        if (IsPressed(ALLEGRO_KEY_D))
        {
            dx = +1;
        }
        if (IsPressed(ALLEGRO_KEY_LSHIFT))
        {
            dy *= 10;
            dx *= 10;
        }
        if (IsPressed(ALLEGRO_MOUSE_KEY_1))
        {
            //cout << "key 1 " << endl;
        }
        if (IsPressed(ALLEGRO_MOUSE_KEY_2))
        {
            //cout << "key 2" << endl;
        }
        humanRobot_.Process(dx, dy);
    }

    virtual void Draw()
    {
        ScreenSaver::Instance().Draw();
        humanRobot_.Draw();
    }

    virtual void OnKeyDown(const ALLEGRO_KEYBOARD_EVENT& keyboard)
    {
        switch (keyboard.keycode)
        {
        case ALLEGRO_KEY_TILDE:
            ScreenSaver::Instance().Reset();
            break;
        case ALLEGRO_KEY_ESCAPE:
            Exit();
            break;
        case ALLEGRO_KEY_SPACE:
            humanRobot_.Fire();
            break;
        }
    }

    virtual void OnMouseMove(const ALLEGRO_MOUSE_EVENT& mouse)
    {
        //cout << mouse.x << " " << mouse.y << endl;
        humanRobot_.SetTarget(mouse.x, mouse.y);
        humanRobot_.TurnGunToTarget();
    }
};

int main(int argc, char** argv)
{
    srand(time(0));

    AllegroApp app;
    if (!app.Init(SCREEN_W, SCREEN_H, FPS))
    {
        return 1;
    }
    app.Run();

    // cin.get();
    return 0;
}
