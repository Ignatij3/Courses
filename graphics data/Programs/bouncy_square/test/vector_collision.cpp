#include <cmath>
#include <iostream>

const double PI = 3.14159265;

class Point {
  public:
    int x, y;

    Point() :
        x(0), y(0) { }

    Point(int x, int y) :
        x(x), y(y) { }

    int XDiff(const Point& rhs) const
    {
        return x - rhs.x;
    }

    int YDiff(const Point& rhs) const
    {
        return y - rhs.y;
    }

    bool operator==(const Point& rhs) const
    {
        return (x == rhs.x && y == rhs.y);
    }
};

// vector defined with starting point and vector going from that point
class Vector {
  private:
    Point a;
    Point b;

  public:
    Vector() { }
    Vector(Point point, Point vector) :
        a(point), b(vector) { }

    Vector(int x1, int y1, int x2, int y2) :
        Vector(Point(x1, y1), Point(x2, y2)) { }

    void Cross(const Vector& lineb) const
    {
        Vector ac(a, lineb.a);
        Vector ad(a, lineb.b);

        int cross1 = *this ^ ac;
        int cross2 = *this ^ ad;
        printf("1) prod: %d %d\n", cross1, cross2);

        Vector ca(lineb.a, a);
        Vector cb(lineb.a, b);

        cross1 = lineb ^ ca;
        cross2 = lineb ^ cb;
        printf("2) prod: %d %d\n\n", cross1, cross2);
    }

    int operator^(Vector& rhs) const
    {
        int x1 = b.XDiff(a);
        int y1 = b.YDiff(a);
        int x2 = rhs.b.XDiff(rhs.a);
        int y2 = rhs.b.YDiff(rhs.a);

        return x1 * y2 - x2 * y1;
    }
};

int main()
{
    Point A(3, 3);
    Point B(9, 7);
    Point C(3, 6);
    Point D(5, 5);
    Point E(6, 5);
    Point F(10, 3);
    Point G(12, 9);
    Point I(6, 3);

    Vector ab(A, B);
    Vector cf(C, F);
    Vector fc(F, C);
    Vector ce(C, E);
    Vector ec(E, C);
    Vector eg(E, G);
    Vector de(D, E);
    Vector ed(E, D);
    Vector cd(C, D);
    Vector dc(D, C);
    Vector ae(A, E);
    Vector ac(A, C);
    Vector ca(C, A);
    Vector ei(E, I);
    Vector ie(I, E);
    Vector ef(E, F);
    Vector fe(F, E);
    Vector bf(B, F);
    Vector fb(F, B);
    Vector ig(I, G);
    Vector gi(G, I);
    Vector bg(B, G);
    Vector gb(G, B);

    Vector left_wall(0, 0, 0, 15);
    Vector bottom_wall(0, 0, 15, 0);
    Vector top_wall(0, 15, 15, 15);
    Vector right_wall(15, 0, 15, 15);

    printf("For ab and cf (cross):\n");
    ab.Cross(cf);
    printf("For ab and fc (cross):\n");
    ab.Cross(fc);
    printf("For ab and ce (touch):\n");
    ab.Cross(ce);
    printf("For ab and ec (touch):\n");
    ab.Cross(ec);
    printf("For ab and cd (don't cross):\n");
    ab.Cross(cd);
    printf("For ab and dc (don't cross):\n");
    ab.Cross(dc);
    printf("For ab and eg (partially lies on):\n");
    ab.Cross(eg);
    printf("For ab and de (touch):\n");
    ab.Cross(de);
    printf("For ab and ed (touch):\n");
    ab.Cross(ed);
    printf("For ab and ab (completely lies on):\n");
    ab.Cross(ab);
    printf("For ab and ae (completely lies on):\n");
    ab.Cross(ae);
    printf("For ab and ac (touch):\n");
    ab.Cross(ac);
    printf("For ab and ca (touch):\n");
    ab.Cross(ca);
    printf("For ab and ie (touch):\n");
    ab.Cross(ie);
    printf("For ab and ei (touch):\n");
    ab.Cross(ei);
    printf("For ab and ig (don't cross):\n");
    ab.Cross(ig);
    printf("For ab and gi (don't cross):\n");
    ab.Cross(gi);
    printf("For ab and fe (touch):\n");
    ab.Cross(fe);
    printf("For ab and ef (touch):\n");
    ab.Cross(ef);
    printf("For ab and bf (touch):\n");
    ab.Cross(bf);
    printf("For ab and fb (touch):\n");
    ab.Cross(fb);
    printf("For ab and bg (touch (on line)):\n");
    ab.Cross(bg);
    printf("For ab and gb (touch (on line)):\n");
    ab.Cross(gb);

    printf("For ab and left_wall (don't touch):\n");
    ab.Cross(left_wall);
    printf("For ab and bottom_wall (don't touch):\n");
    ab.Cross(bottom_wall);
    printf("For ab and top_wall (don't touch):\n");
    ab.Cross(top_wall);
    printf("For ab and right_wall (don't touch):\n");
    ab.Cross(right_wall);

    return 0;
}