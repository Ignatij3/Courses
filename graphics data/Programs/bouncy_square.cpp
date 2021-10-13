#include "AllegroUtil.hpp"

#include <cmath>
#include <cstdio>
#include <iostream>
#include <vector>

const int FPS      = 200;
const int SCREEN_W = 1200;
const int SCREEN_H = 720;

namespace window
{
#define PI 3.14159265

    class Point {
      public:
        int x, y;

        Point() :
            x(0), y(0) { }

        Point(int x, int y) :
            x(x), y(y) { }

        int xdiff(const Point& rhs) const
        {
            return this->x - rhs.x;
        }

        int ydiff(const Point& rhs) const
        {
            return this->y - rhs.y;
        }
    };

    //line defined with starting point and vector, going from that point
    class Line {
      private:
        double angle;

        void setAngle()
        {
            this->angle = atan2(a.ydiff(b), a.xdiff(b)) * 180 / PI;
            this->angle = (this->angle < 0) ? 360 + this->angle : this->angle;
        }

      public:
        Point a;
        Point b;
        Line() { setCoordinates(0, 0, 0, 0); }

        Line(Point point, Point vector) :
            a(point), b(vector) { setAngle(); }

        void setCoordinates(int x1, int y1, int x2, int y2)
        {
            this->a.x = x1;
            this->a.y = y1;
            this->b.x = x2;
            this->b.y = y2;

            setAngle();
        }

        double getAngle() const { return this->angle; }

        int operator^(Line& rhs) const
        {
            int x1 = this->b.xdiff(this->a);
            int y1 = this->b.ydiff(this->a);
            int x2 = rhs.b.xdiff(rhs.a);
            int y2 = rhs.b.ydiff(rhs.a);

            return x1 * y2 - x2 * y1;
        }

        bool cross(const Line& lineb) const
        {
            auto sign = [](int x) -> char { return (x >= 0 ? '+' : '-'); };

            Line ac(this->a, lineb.a);
            Line ad(this->a, lineb.b);
            int cross1 = *this ^ ac;
            int cross2 = *this ^ ad;

            if (sign(cross1) == sign(cross2) || cross1 == 0 || cross2 == 0)
                return false;

            Line ca(lineb.a, this->a);
            Line da(lineb.a, this->b);
            cross1 = lineb ^ ca;
            cross2 = lineb ^ da;

            return !(sign(cross1) == sign(cross2) || cross1 == 0 || cross2 == 0);
        }
    };

    const Line walls[4] = {
        Line(Point(0, 0), Point(SCREEN_W, 0)),               //top wall
        Line(Point(SCREEN_W, 0), Point(SCREEN_W, SCREEN_H)), //right wall
        Line(Point(0, 0), Point(0, SCREEN_H)),               //left wall
        Line(Point(0, SCREEN_H), Point(SCREEN_W, SCREEN_H))  //bottom wall
    };

    class Square {
      private:
        double x, y; //center coordinates
        double halfSide;
        double angle;
        std::pair<double, double> direction; //angle of square direction in degrees

        Line sides[4];

        double leftX() const { return x - halfSide; }
        double upperY() const { return y - halfSide; }
        double rightX() const { return x + halfSide; }
        double lowerY() const { return y + halfSide; }

        void setSides()
        {
            double lx = this->leftX();
            double rx = this->rightX();
            double uy = this->upperY();
            double ly = this->lowerY();

            this->sides[0].setCoordinates(lx, uy, this->halfSide * 2, 0); //top side
            this->sides[1].setCoordinates(rx, uy, 0, this->halfSide * 2); //right side
            this->sides[2].setCoordinates(lx, uy, 0, this->halfSide * 2); //left side
            this->sides[3].setCoordinates(lx, ly, this->halfSide * 2, 0); //bottom side
        }

        const Line* getCollidedVector() const
        {
            for (int i = 0; i < 4; ++i)
            {
                for (int j = 0; j < 4; ++j)
                {
                    if (sides[i].cross(walls[j]))
                        return &walls[j];
                }
            }

            return nullptr;
        }

        void reflect(double otherVectorAngle)
        {
            double newAngle = -360 + (2 * otherVectorAngle) - this->angle;
            newAngle        = (newAngle < 0) ? 360 + newAngle : newAngle;
            setDirection(newAngle);
        }

      public:
        static Square initFromStdin()
        {
            double x, y;
            double side;
            double angle;

            printf("Enter x-coordinate of square's centre: ");
            scanf("%f\n", x);

            printf("Enter y-coordinate of square's centre: ");
            scanf("%f\n", y);

            printf("Enter square's side: ");
            scanf("%f\n", side);

            printf("Enter square's movement angle: ");
            scanf("%f\n", angle);

            return Square(x, y, side, angle);
        }

        Square(double xcoord, double ycoord, double squareSide, double alpha) :
            halfSide((squareSide < 0) ? 0 : squareSide / 2),
            x(xcoord),
            y(ycoord)
        {
            x = (leftX() < 0) ? halfSide
                              : ((x + halfSide > SCREEN_W) ? SCREEN_W - halfSide : x);
            y = (upperY() < 0) ? halfSide
                               : ((y + halfSide > SCREEN_H) ? SCREEN_H - halfSide : y);

            setDirection(alpha);
            setSides();
        }

        //setDirection sets rectangle's direction in degrees, where 0 points right and goes anticlockwise
        //if alpha is smaller than 0 or greater than 360, direction is set to 0
        void setDirection(double alpha)
        {
            this->angle = (alpha < 0 || alpha > 360) ? 0 : alpha;

            this->direction.first  = cos(angle * PI / 180); //converting degrees to radians
            this->direction.second = sin(angle * PI / 180); //converting degrees to radians
        }

        void next()
        {
            if (const Line* vec = getCollidedVector(); vec != nullptr)
                reflect(vec->getAngle());

            this->x += direction.first;
            this->y -= direction.second;

            setSides();
        }

        void draw(ALLEGRO_COLOR color)
        {
            double lx = this->leftX();
            double rx = this->rightX();
            double uy = this->upperY();
            double ly = this->lowerY();
            al_draw_filled_rectangle(lx, uy, rx, ly, color);
        }
    };
}

std::vector<window::Square> figures;

void nextFrame()
{
    for (auto& fig : figures)
    {
        fig.next();
    }
}

void draw()
{
    ALLEGRO_COLOR RED   = al_map_rgb(255, 0, 0);
    ALLEGRO_COLOR WHITE = al_map_rgb(255, 255, 255);

    al_clear_to_color(WHITE);
    for (auto& fig : figures)
    {
        fig.draw(RED);
    }
}

int main()
{
    //window::Square sqr1 = window::Square::initFromStdin();
    window::Square sqr1(100, 500, 40, 135);
    window::Square sqr2(200, 200, 75, 330);

    figures.push_back(sqr1);
    figures.push_back(sqr2);

    if (!InitAllegro(SCREEN_W, SCREEN_H, FPS))
    {
        DestroyAllegro();
        return 1;
    }

    RunAllegro(&nextFrame, &draw);
    ExitAllegro();

    return 0;
}