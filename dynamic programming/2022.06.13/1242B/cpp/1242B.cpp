#include <bitset>
#include <cstdio>
#include <iostream>
#include <vector>

typedef std::vector<std::bitset<100000>> wgraph;

void traverseComponent(wgraph*, int, int);

int main()
{
    int n, m;
    scanf("%d %d", &n, &m);
    getchar();

    wgraph graph(n, std::bitset<100000>());
    for (int i = 0; i < n; ++i)
    {
        graph[i][i] = true;
    }

    int a, b;
    for (int i = 0; i < m; ++i)
    {
        scanf("%d %d", &a, &b);
        getchar();
        --a, --b;
        graph[a][b] = true;
        graph[b][a] = true;
    }

    int components = 0;

    int nodes = graph.size();

    for (int i = 0; i < nodes; ++i)
    {
        bool noZeros = true;
        for (int j = 0; j < nodes; ++j)
        {
            if (!graph[i][j])
            {
                noZeros = false;
            }
        }
        if (noZeros)
        {
            ++components;
        }
    }

    for (int i = 0; i < nodes; ++i)
    {
        for (int j = i + 1; j < nodes; ++j)
        {
            if (!graph[i][j])
            {
                traverseComponent(&graph, i, j);
                components++;
            }
        }
    }

    std::cout << --components;
    return 0;
}

void traverseComponent(wgraph* graph, int a, int b)
{
    (*graph)[a][b] = true;
    (*graph)[b][a] = true;
    for (int i = 0; i < (*graph).size(); ++i)
    {
        if (!(*graph)[a][i])
        {
            traverseComponent(graph, a, i);
        }
        if (!(*graph)[b][i])
        {
            traverseComponent(graph, b, i);
        }
    }
}