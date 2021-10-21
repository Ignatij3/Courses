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

        if (alTimer = al_create_timer(1.0 / fps); !alTimer)
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

    void Window::RunAllegro(FpsCallback fpsCallback, DrawCallback drawCallback)
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

            // printf("%s ", redraw ? "true" : "false");
            // printf("%s\n", al_is_event_queue_empty(alEventQueue) ? "true" : "false");
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

        if (fps <= 0 || window_width <= 0 || window_height <= 0)
            throw InitFailure("Window error: Invalid window initialization arguments passed\n");

        printf("width, height: %f, %f\n", double(window_width / 2), double(window_height / 2));
        window_frame       = shape::Rectangle(shape::Movement::STATIC, window_width / 2, window_height / 2, window_width, window_height, 0);
        window_frame.color = LIGHTGRAY;
        objects.insert(objects.end(), &window_frame);

        try
        {
            InitAllegro();

        } catch (const InitFailure& e)
        {
            DestroyAllegro();
            throw InitFailure(e.what());
        }

        initialized = true;
    }

    // starts main event loop
    void Window::Run(FpsCallback fpsCallback, DrawCallback drawCallback)
    {
        RunAllegro(fpsCallback, drawCallback);
        exit = true;
    }

    void Window::AddObject(shape::Shape& figure)
    {
        if (!*this)
            throw InitFailure("Window error: Window not initialised\n");
        objects.insert(objects.end(), &figure);
    }

    void Window::SetColor(shape::Shape& figure, ALLEGRO_COLOR col) noexcept
    {
        figure.color = col;
    }

    void Window::AddObject(shape::Shape& figure, ALLEGRO_COLOR col)
    {
        SetColor(figure, col);
        AddObject(figure);
    }

    // check for collisions between objects
    void Window::ManageCollisions() const noexcept
    {
        if (objects.size() > 1)
            for (int i = 0; i < objects.size(); ++i)
                for (int j = i; j < objects.size(); ++j)
                {
                    auto res = objects[i]->CollideWith(objects[j]);
                    if (res.first)
                    {
                        objects[i]->Reflect(res.second.second.getAngle());
                        objects[j]->Reflect(res.second.first.getAngle());
                        return;
                    }
                }
    }

    void Window::MoveAll() noexcept
    {
        for (auto obj = objects.begin(); obj != objects.end(); ++obj)
        {
            (*obj)->Move();
        }
    }

    void Window::DrawAll() const noexcept
    {
        for (auto obj = objects.begin(); obj != objects.end(); ++obj)
            (*obj)->Draw();
    }

    Window::operator bool() const
    {
        return initialized;
    }
}
