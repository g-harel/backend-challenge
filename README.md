# Backend Intern Challenge - Summer 2018

## Problem

Different merchants have different products. Menus are a great way to present these products in an intuitive way to customers.

To address this we decided to build a system that allows Menus validation. Our system should be able to aggregate the provided Menus and identify if any of them contain cyclical references.

Nodes are expressed as objects in a JSON array with the following fields:

| Fields      | Possible Values     | Optional Values | Description                                |
| ----------- | ------------------- | --------------- | ------------------------------------------ |
| `id`        | `"number"`          |                 | Uniquely identifies each node of the Menu. |
| `data`      | `"string"`          |                 | Specifies the data that the node holds.    |
| `parent_id` | `"number"`          | ✓               | Specifies the parent of the node.          |
| `child_ids` | Array of `"number"` |                 | Specifies the children of the node.        |

### API Example Response

```
{
  "menus":[
    {
      "id":1,
      "data":"House",
      "child_ids":[3]
    },
    {
      "id":2,
      "data":"Company",
      "child_ids":[4]
    },
    {
      "id":3,
      "data":"Kitchen",
      "parent_id":1,
      "child_ids":[5]
    },
    {
      "id":4,
      "data":"Meeting Room",
      "parent_id":2,
      "child_ids":[6]
    },
    {
      "id":5,
      "data":"Sink",
      "parent_id":3,
      "child_ids":[1]
    },
    {
      "id":6,
      "data":"Chair",
      "parent_id":4,
      "child_ids":[]
    }
  ],
  "pagination":{
    "current_page":1,
    "per_page":5,
    "total":19
  }
}
```

## API Response

The response will contain the followings keys:

- `menus` - Array containing nodes that together become Menus.
- `pagination` - An object containing the `current_page`, `per_page` and `total keys`

## Instructions

**Candidates can use any programming language of their choice.**

Obtain a list of Nodes from the API and build the appropriate Menus. Cycles should be identified and the offending Menus should be marked as invalid. The max depth of a Menu is **4**.

The output for the given example should be as follows:

```
{
  "valid_menus": [
    { "root_id": 2, "children": [4, 6] },
  ],
  "invalid_menus": [
    { "root_id": 1, "children": [1, 3, 5] }
  ]
}
```

The output is expected to be in JSON and to contain the following keys:

- `valid_menus` - Array containing information about each valid menu:
    - `root_id` - The id of the menu root
    - `children` - An Array containing all the children belonging to the menu
- `invalid_menus` - Array containing information about each invalid menu:
    - `root_id` - The id of the menu root
    - `children` - An Array containing all the children belonging to the menu

The API endpoint can be found at:

https://backend-challenge-summer-2018.herokuapp.com/challenges.json?id=1&page=1

This will return the first page of nodes. You can obtain subsequent pages by incrementing the page query parameter.

## Extra Challenge

The same rules of the original challenge applies.

The API endpoint can be found at:

https://backend-challenge-summer-2018.herokuapp.com/challenges.json?id=2&page=1

This will return the first page of nodes. You can obtain subsequent pages by incrementing the page query parameter.
