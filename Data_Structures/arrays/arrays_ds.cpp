#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
using namespace std;


int main() {
    /* Enter your code here. Read input from STDIN. Print output to STDOUT */  
    vector<int> arr;
    int n;
    int input;
    
    cin >> n;
    
    while(n--){
        cin >> input;
        arr.push_back(input);
    }
    
    reverse(arr.begin(), arr.end());
    
    for(vector<int>::iterator it = arr.begin();
        it != arr.end();
        it++){
        cout << *it << " ";
    }
    
    return 0;
}
