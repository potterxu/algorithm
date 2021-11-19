#include "union_find.hpp"

void UnionFind::add(int v)
{
    if (m_nodeMap.find(v) != m_nodeMap.end()) {
        printf("Node %d already existed", v);
        return;
    }
    m_nodeMap[v] = std::make_shared<Node>(v);
    m_nodeMap[v]->root = m_nodeMap[v];
}

int UnionFind::find(int v)
{
    if (m_nodeMap.find(v) == m_nodeMap.end()) {
        printf("Node %d does not found", v);
        return -1;
    }
    return _find(m_nodeMap[v])->value;
}

void UnionFind::group(int a, int b)
{
    if (m_nodeMap.find(a) == m_nodeMap.end()) {
        add(a);
    }
    if (m_nodeMap.find(b) == m_nodeMap.end()) {
        add(b);
    }
    return _group(m_nodeMap[a], m_nodeMap[b]);
}

std::shared_ptr<UnionFind::Node> UnionFind::_find(std::shared_ptr<UnionFind::Node> node)
{
    if (node == node->root) {
        return node;
    }
    node->root =  _find(node->root);
    return node->root;
}

void UnionFind::_group(std::shared_ptr<UnionFind::Node> a, std::shared_ptr<UnionFind::Node> b)
{
    auto small = _find(a);
    auto large = _find(b);

    if (small->rank > large->rank) {
        std::swap(small, large);
    }

    small->root = large;

    if (small->rank == large->rank) {
        large->rank++;
    }
}