#include <iostream>
#include <vector>
using namespace std;

void static print_arr(vector<int> *arr){
    for(vector<int>::iterator it = arr->begin();
        it != arr->end();
        it++){
        
        cout << *it << " ";
    }
    cout << endl;
}

vector<int> static quick_sort(vector<int> arr, int arr_size) {
    vector<int> left, right;
    vector<int>::iterator it = arr.begin();
    int pivot;
    
    if(arr_size <= 1){
        return arr;
    }
    
    pivot = *it;
    it++;
    for(int i=1; i<arr.size(); i++, it++){
        if(*it < pivot){
            left.push_back(*it);
        }
        else{
            right.push_back(*it);
        }
    }
    
    left = quick_sort(left, left.size());
    right = quick_sort(right, right.size());
    left.push_back(pivot);
    
    left.reserve(left.size() + right.size());
    left.insert(left.end(), right.begin(), right.end());
    
    print_arr(&left);
    return left;
}

int main(void)
{
    int n;
    vector<int> arr;
    cin >> n;

    for(int i = 0; i < n; i++) {
        int tmp_input;
        cin >> tmp_input;
        arr.push_back(tmp_input);
    }

    quick_sort(arr, arr.size());

    return 0;
}

