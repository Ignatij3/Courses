#include "screen.hpp"

#include <iostream>
#include <string>
#include <vector>

namespace screen
{
    shape::Vector walls[4]; //TODO write copy-contructor

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

        walls[0].SetCoordinates(0, 0, win_width, 0);                   //top wall
        walls[1].SetCoordinates(win_width, 0, win_width, win_height);  //right wall
        walls[2].SetCoordinates(0, 0, 0, win_height);                  //left wall
        walls[3].SetCoordinates(0, win_height, win_width, win_height); //bottom wall

        printf("vec x1, y1: %f %f\n", walls[1].a.x, walls[1].a.y);
        printf("vec x2, y2: %f %f\n", walls[1].b.x, walls[1].b.y);
        objects.first.push_back(&walls[0]);
        objects.first.push_back(&walls[1]);
        objects.first.push_back(&walls[2]);
        objects.first.push_back(&walls[3]);
    }

    //starts main event loop
    void Window::Run(FpsCallback fpsCallback, DrawCallback drawCallback)
    {
        RunAllegro(fpsCallback, drawCallback);
        ExitAllegro();
    }

    void Window::AddObject(shape::Vector* vec)
    {
        objects.first.insert(objects.first.end(), vec);
    }

    void Window::AddObject(shape::Shape* figure)
    {
        objects.second.insert(objects.second.end(), figure);
    }

    void Window::SetColor(shape::Shape* figure, ALLEGRO_COLOR col)
    {
        figure->color = col;
    }

    void Window::AddObject(shape::Shape* figure, ALLEGRO_COLOR col)
    {
        SetColor(figure, col);
        AddObject(figure);
    }

    void Window::ManageCollisions() const
    {

        for (auto object1 = objects.second.begin(); object1 != objects.second.end(); ++object1)
        {
            //check for collisions between objects
            if (objects.second.size() > 1)
            {
                for (auto object2 = ++object1; object2 != objects.second.end(); ++object2)
                {
                    std::pair<bool, const shape::Vector> res = (*object1)->getVectorIfCollide(*object2);
                    if (res.first)
                    {
                        (*object1)->Reflect(res.second.getAngle());
                        (*object1)->Move();
                        (*object2)->Move();
                    }
                }
            }

            //check for collisions between object and a vector
            for (auto vector = objects.first.begin(); vector != objects.first.end(); ++vector)
            {
                std::pair<bool, const shape::Vector> res = (*object1)->getVectorIfCollide(*vector);

                printf("vec x1, y1: %f %f\n", walls[1].a.x, walls[1].a.y);
                printf("vec x2, y2: %f %f\n", walls[1].b.x, walls[1].b.y);

                if (res.first)
                {
                    (*object1)->Reflect(res.second.getAngle());
                    (*object1)->Move();
                }
            }
        }
    }

    void Window::MoveAll()
    {
        for (auto obj = objects.second.begin(); obj != objects.second.end(); ++obj)
            (*obj)->Move();
    }

    void Window::DrawAll() const
    {
        for (auto obj = objects.second.begin(); obj != objects.second.end(); ++obj)
            (*obj)->Draw();
    }

}