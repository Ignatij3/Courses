#include "screen.hpp"

#include <iostream>
#include <string>
#include <vector>

namespace screen
{
    std::pair<double, double> ConvertToNormalCoords(double x, double y)
    {
        return std::make_pair(x, Window::window_height - y);
    }

    shape::Point ConvertToNormalCoords(shape::Point a)
    {
        std::pair<double, double> pt = ConvertToNormalCoords(a.x, a.y);
        return shape::Point(pt.first, pt.second);
    }

    Window::BadInit::BadInit(const std::string& msg) :
        message(msg) { }

    const char* Window::BadInit::what() const throw()
    {
        return message.c_str();
    }

    int Window::fps;
    int Window::window_width;
    int Window::window_height;

    Window::Window(int fps_, int win_width, int win_height)
    {
        fps           = fps_;
        window_width  = win_width;
        window_height = win_height;

        if (fps <= 0 || window_width <= 0 || window_height <= 0) throw BadInit("Window error: Invalid window initialization arguments passed\n");

        if (!InitAllegro(window_width, window_height, fps))
        {
            DestroyAllegro();
            throw BadInit("Allegro init error: Allegro initialization fail\n");
        }

        walls[0].SetVectors(0, 0, win_width, 0);          // top wall
        walls[1].SetVectors(win_width, 0, 0, win_height); // right wall
        walls[2].SetVectors(0, 0, 0, win_height);         // left wall
        walls[3].SetVectors(0, win_height, win_width, 0); // bottom wall

        walls[0].setAngle();
        walls[1].setAngle();
        walls[2].setAngle();
        walls[3].setAngle();

        objects.first.push_back(walls[0]);
        objects.first.push_back(walls[1]);
        objects.first.push_back(walls[2]);
        objects.first.push_back(walls[3]);
    }

    // starts main event loop
    void Window::Run(FpsCallback fpsCallback, DrawCallback drawCallback)
    {
        RunAllegro(fpsCallback, drawCallback);
        ExitAllegro();
    }

    void Window::AddObject(shape::Vector vec)
    {
        objects.first.insert(objects.first.end(), vec);
    }

    void Window::AddObject(shape::Shape& figure)
    {
        objects.second.insert(objects.second.end(), &figure);
    }

    void Window::SetColor(shape::Shape& figure, ALLEGRO_COLOR col)
    {
        figure.color = col;
    }

    void Window::AddObject(shape::Shape& figure, ALLEGRO_COLOR col)
    {
        SetColor(figure, col);
        AddObject(figure);
    }

    void Window::ManageCollisions() const
    {
        for (int i = 0; i < objects.second.size(); ++i)
        {
            // check for collisions between objects
            if (objects.second.size() > 1)
            {
                for (int j = 0; j < objects.second.size(); ++j)
                {
                    auto res = objects.second[i]->CollideWith(objects.second[j]);
                    if (res.first)
                    {
                        objects.second[i]->Reflect(res.second.second.getAngle());
                        objects.second[j]->Reflect(res.second.first.getAngle());
                        return;
                    }
                }
            }

            // check for collisions between object and a vector
            for (int vec = 0; vec < objects.first.size(); ++vec)
            {
                std::pair<bool, const shape::Vector> res = objects.second[i]->CollideWith(&objects.first[vec]);

                if (res.first)
                {
                    objects.second[i]->Reflect(res.second.getAngle());
                    return;
                }
            }
        }
    }

    void Window::MoveAll()
    {
        for (auto obj = objects.second.begin(); obj != objects.second.end(); ++obj)
        {
            (*obj)->Move();
        }
    }

    void Window::DrawAll() const
    {
        for (auto obj = objects.second.begin(); obj != objects.second.end(); ++obj)
            (*obj)->Draw();
    }
}
