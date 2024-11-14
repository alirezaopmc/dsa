#include <iostream>>

typedef long long ll;

using namespace std;

const int N = 300000;
ll dp[N][2];
ll a[N];
int parent[N];

void test_case()
{
    int n;
    cin >> n;

    // Attack Values
    for (int i = 0; i < n; i++)
    {
        cin >> a[i];
    }

    // Constructing the tree
    for (int i = 0; i < n - 1; i++)
    {
        int x, y;
        cin >> x >> y;
        x--;
        y--;
        parent[y] = x;
    }

    // Initializing the dp
    for (int i = 0; i < n; i++)
    {
        dp[i][0] = 0;
        dp[i][1] = a[i];
    }
}

int main()
{
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
}