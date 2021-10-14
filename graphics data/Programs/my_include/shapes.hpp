#ifndef GRAPHICS_SHAPES_H_
#define GRAPHICS_SHAPES_H_

#include "AllegroUtil.hpp"

#include <vector>

namespace shape
{
    const double PI = 3.14159265;

    template <class T>
    class Point {
      public:
        T x, y;

        Point() :
            x(0), y(0) { }

        Point(T x, T y) :
            x(x), y(y) { }

        T XDiff(const Point& rhs) const
        {
            return this->x - rhs.x;
        }

        T YDiff(const Point& rhs) const
        {
            return this->y - rhs.y;
        }
    };

    // vector defined with starting point and vector going from that point
    class Vector {
      private:
        double angle;

        void setAngle();

      public:
        Point<double> a;
        Point<double> b;
        ALLEGRO_COLOR color;

        Vector();
        Vector(Point<double> point, Point<double> vector);
        Vector(double x1, double y1, double x2, double y2);

        void SetCoordinates(Point<double> point, Point<double> vector);
        void SetCoordinates(double x1, double y1, double x2, double y2);
        double getAngle() const;
        bool Cross(const Vector& lineb) const;
        double Magnitude();
        int operator^(Vector& rhs) const;
        Vector& operator=(Vector&& rhs);
    };

    // in circle, check for collision with vector pointing in move direction
    class Shape {
      protected:
        double angle;
        std::pair<double, double> direction;

        virtual double LeftmostX() const  = 0;
        virtual double UppermostY() const = 0;
        virtual double RightmostX() const = 0;
        virtual double LowermostY() const = 0;
        virtual void SetSides()           = 0;

      public:
        Vector sides[0];
        Point<double> centre;
        ALLEGRO_COLOR color = al_map_rgb(0, 0, 0);

        Shape(Point<double> centreCoords, double alpha);

        virtual const int sideAmount() const = 0;
        virtual Shape* InitFromStdin() const = 0;
        virtual void Move()                  = 0;
        virtual void Draw() const            = 0;
        void Reflect(double otherVectorAngle);

        // setDirection sets shape's direction in degrees, where 0 points right and goes anticlockwise
        // if alpha is smaller than 0 or greater than 360, direction is set to 0
        void SetDirection(double alpha);
        std::pair<bool, const Vector*> getVectorIfCollide(const Shape* other) const;
        std::pair<bool, const Vector*> getVectorIfCollide(const Vector* other_vector) const;
    };

    class Rectangle : public Shape {
    };

    class Square : public Shape {
      private:
        double edge; // distance from square's centre to edge

        double LeftmostX() const override;
        double UppermostY() const override;
        double RightmostX() const override;
        double LowermostY() const override;
        void SetSides() override;

      public:
        Vector sides[4];
        Square(Point<double> centreCoords, double side, double alpha);
        Square(double centreX, double centreY, double side, double alpha);

        const int sideAmount() const override;
        Shape* InitFromStdin() const override;
        void Move() override;
        void Draw() const override;
    };
}

#endif // !GRAPHICS_SHAPES_H_