# Coding Rule

## Controller method

| Controller Method | HTTP Method |                 Description                 |
| ----------------- | ----------- | ------------------------------------------- |
| Index             | GET         | Displays a listing of the resource          |
| Store             | POST        | Stores a newly created resource in storage  |
| Show              | GET         | Displays the specified resource             |
| Update            | PUT/PATCH   | Updates the specified resource in storage   |
| Destroy           | DELETE      | Removes the specified resource from storage |

## Repository method

| Repository Method |                     Description                      |
| ----------------- | ---------------------------------------------------- |
| FindByXX          | Returns the entity identified by the given XX        |
| FindAll           | Returns all entities                                 |
| FindAllXX         | Returns all XX entities                              |
| Save              | Saves the given entity                               |
| SaveByXX          | Saves the given entity identified by the given XX    |
| DeleteByXX        | Deletes the entity identified by the given XX        |
| Count             | Returns the number of entities                       |
| ExistsBy          | Indicates whether an entity with the given ID exists |