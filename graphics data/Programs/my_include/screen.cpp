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

    Window::InitFailure::InitFailure(const std::string& msg) :
        message(msg) { }

    const char* Window::InitFailure::what() const throw()
    {
        return message.c_str();
    }

    int Window::fps;
    int Window::window_width;
    int Window::window_height;

    void Window::InitAllegro()
    {
        if (!al_init())
            throw InitFailure("Allegro init error: Failed to initialize allegro\n");

        if (alTimer = al_create_timer(1 / fps); !alTimer)
            throw InitFailure("Allegro init error: Failed to create timer\n");

        if (alDisplay = al_create_display(window_width, window_height); !alDisplay)
            throw InitFailure("Allegro init error: Failed to create display\n");

        if (!al_init_primitives_addon())
            throw InitFailure("Allegro init error: Failed to init addons\n");

        if (alEventQueue = al_create_event_queue(); !alEventQueue)
            throw InitFailure("Allegro init error: Failed to create event queue\n");

        al_register_event_source(alEventQueue, al_get_display_event_source(alDisplay));
        al_register_event_source(alEventQueue, al_get_timer_event_source(alTimer));
    }

    void Window::DestroyAllegro()
    {
        if (alTimer)
            al_destroy_timer(alTimer);

        if (alDisplay)
            al_destroy_display(alDisplay);

        if (alEventQueue)
            al_destroy_event_queue(alEventQueue);
    }

    void Window::RunAllegro(void (*fpsCallback)(), void (*drawCallback)())
    {
        // clear screen
        al_clear_to_color(al_map_rgb(255, 255, 255));
        al_flip_display();

        al_start_timer(alTimer);

        bool redraw = false;
        while (true)
        {
            ALLEGRO_EVENT ev;
            al_wait_for_event(alEventQueue, &ev);

            if (ev.type == ALLEGRO_EVENT_TIMER)
            {
                if (fpsCallback)
                    fpsCallback();
                redraw = true;
            }
            else if (ev.type == ALLEGRO_EVENT_DISPLAY_CLOSE)
                break;

            if (redraw && al_is_event_queue_empty(alEventQueue))
            {
                redraw = false;
                if (drawCallback)
                    drawCallback();

                al_flip_display();
            }

            if (exit)
                break;
        }
    }

    Window::Window(int fps_, int win_width, int win_height)
    {
        fps           = fps_;
        window_width  = win_width;
        window_height = win_height;

        if (fps <= 0 || window_width <= 0 || window_height <= 0) throw InitFailure("Window error: Invalid window initialization arguments passed\n");

        try
        {
            InitAllegro();

        } catch (const InitFailure& e)
        {
            DestroyAllegro();
            throw InitFailure(e.what());
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
    void Window::Run(void (*fpsCallback)(), void (*drawCallback)())
    {
        RunAllegro(fpsCallback, drawCallback);
        exit = true;
    }

    void Window::AddObject(shape::Vector vec) noexcept
    {
        objects.first.insert(objects.first.end(), vec);
    }

    void Window::AddObject(shape::Shape& figure) noexcept
    {
        objects.second.insert(objects.second.end(), &figure);
    }

    void Window::SetColor(shape::Shape& figure, ALLEGRO_COLOR col) noexcept
    {
        figure.color = col;
    }

    void Window::AddObject(shape::Shape& figure, ALLEGRO_COLOR col) noexcept
    {
        SetColor(figure, col);
        AddObject(figure);
    }

    void Window::ManageCollisions() const noexcept
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

    void Window::MoveAll() noexcept
    {
        for (auto obj = objects.second.begin(); obj != objects.second.end(); ++obj)
        {
            (*obj)->Move();
        }
    }

    void Window::DrawAll() const noexcept
    {
        for (auto obj = objects.second.begin(); obj != objects.second.end(); ++obj)
            (*obj)->Draw();
    }
}
