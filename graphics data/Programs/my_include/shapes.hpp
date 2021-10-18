#ifndef GRAPHICS_SHAPES_H_
#define GRAPHICS_SHAPES_H_

#include "AllegroUtil.hpp"

#include <vector>

bool AlmostEqual(double a, double b, double epsilon);

namespace shape
{
    const double PI = 3.14159265;

    class Point {
      public:
        double x, y;

        Point() :
            x(0), y(0) { }

        Point(double x, double y) :
            x(x), y(y) { }

        double XDiff(const Point& rhs) const
        {
            return this->x - rhs.x;
        }

        double YDiff(const Point& rhs) const
        {
            return this->y - rhs.y;
        }

        bool operator==(const Point& rhs) const
        {
            return (AlmostEqual(x, rhs.x, 0.4) && AlmostEqual(y, rhs.y, 0.4));
        }
    };

    // vector defined with starting point and vector going from that point
    class Vector {
      private:
        double angle;

      public:
        ALLEGRO_COLOR color;
        Point a;
        Point b;

        Vector();
        Vector(Point point, Point vector);
        Vector(double x1, double y1, double x2, double y2);
        Vector(const Vector& rhs) noexcept;

        double Slope() const;
        bool LiesBetween(const Vector& outsideVector) const;
        void setAngle();
        void SetVectors(Point point, Point vector);
        void SetVectors(double x1, double y1, double x2, double y2);
        double getAngle() const;
        bool Cross(const Vector& lineb) const;
        double Magnitude();
        double HighestX() const;
        double LowestX() const;
        double HighestY() const;
        double LowestY() const;
        int operator^(Vector& rhs) const;
        Vector& operator=(const Vector& rhs);
        Vector& operator=(Vector&& rhs);
        bool operator==(const Vector& rhs) const;
    };

    // in circle, check for collision with vector pointing in move direction
    class Shape { //брать в расчёт массу и потерю скорости //stationary option
      protected:
        double angle;
        std::vector<Vector> sides;
        Point centre;
        std::pair<double, double> direction;

        virtual double LeftmostX() const  = 0;
        virtual double UppermostY() const = 0;
        virtual double RightmostX() const = 0;
        virtual double LowermostY() const = 0;
        virtual void SetSides()           = 0;
        virtual void SetAngleSides()      = 0;
        std::pair<bool, const Vector> LiesOnLine(const std::vector<Vector>& sides, const Point& angle) const;
        std::pair<const Vector, const Vector> FindSidesToReflect(std::vector<Vector>& shapeSides, std::vector<Vector>& otherShapeSides, int sideIndex, int otherSideIndex) const;

      public:
        ALLEGRO_COLOR color = al_map_rgb(0, 0, 0);

        Shape(Point centreCoords, double width, double height, double alpha);

        virtual std::vector<Vector> GetSides() const = 0;
        virtual const int sideAmount() const         = 0;
        virtual void Move()                          = 0;
        virtual void Draw() const                    = 0;
        void Reflect(double otherVectorAngle);

        // setDirection sets shape's direction in degrees, where 0 points right and goes anticlockwise
        // if alpha is smaller than 0 or greater than 360, direction is set to 0
        void SetDirection(double alpha);
        std::pair<bool, std::pair<const Vector, const Vector>> CollideWith(const Shape* other) const;
        std::pair<bool, const Vector> CollideWith(const Vector* other_vector) const;
    };

    class Rectangle : public Shape {
    };

    class Square : public Shape {
      private:
        std::vector<Vector> sides;
        double edge; // distance from square's centre to edge

        double LeftmostX() const override;
        double UppermostY() const override;
        double RightmostX() const override;
        double LowermostY() const override;
        void SetSides() override;
        void SetAngleSides() override;

      public:
        Square(Point centreCoords, double side, double alpha);
        Square(double centreX, double centreY, double side, double alpha);

        std::vector<Vector> GetSides() const override;
        const int sideAmount() const override;
        Square static InitFromStdin();
        void Move() override;
        void Draw() const override;
    };
}

#endif // !GRAPHICS_SHAPES_H_