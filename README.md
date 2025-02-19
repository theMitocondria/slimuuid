# SlimUUID ✨

SlimUUID is a 18 character long Unique Idnetifier . For detailed information you can give this file a read . 

## Table of Contents

1. [UUIDs and slimUUID](#uuids-and-slimuuid)  
2. [Components of SlimUUID](#components-of-slimuuid)  
3. [Technical Details](#technical-details)  
4. [Installation](#installation)  
5. [Usage](#usage)  
6. [Examples](#examples)  
7. [Contributing](#contributing)  
8. [License](#license)

## UUIDs and slimUUID

Unique Identifiers are used in Databases for giving a unique identifiers to each entry . We can have different types varieing from integers , ObjectID(For MongoID) , strings . 

But these unique identifiers must have some properties that are needed for different purposes :

1. Unpredictablity because they used in urls 
2. Collision Probability to be autronomically low because used as Primary key
3. Short length for storage , indexing and searching
4. TimeStamped , can be and can be not . But is beneficial to have it timestamped .

So while keeping in mind all of these , We shall learn how we made slimUUID , how you can use it .

**Unpredictability** : As you can also use integers for Unique Identifiers and you would save a lot of space but it is not good where you are sending these IDs in url to customer .
- Solution : SlimUUID is using its last 8-10 characters for MAC Address , IP Address , xxHash Seed Hashing with each nanosecond of time . So predictability of this uuid is astronomically low . 


**Collision Probability** : When generating a UUID one thing that is a must is its neglible in collision probabilty . By collision probabilty I mean the probability of same id being generated after how many generation and at what percentage of repeatition .
- Solution : SlimUUID uses precision till NanoSeconds , For reference a general Rand.Intn() function take 1ms which is 1000 nanoseconds . So while using nanoseconds , we keep in mind the collison which can occur when multiple systems are involved or same systems threads are used , which btw is next to impossible considering current hardware contraints because even switching context take more than 10ns . So we assign each ID a unique Address , kindoff Global Thread counter . So getting hit by astroid is much more probable then getting hit by slimUUID collison . 
So for proof we got it calculated , [For Exact reference Go to this Blog](https://dev.to/mitocondria/slimuuid-the-compact-memory-efficient-alternative-to-standard-uuids-2dak)


**Space and Time Complexity** : So Each uuid takes some time to be generated , needs a space to be stored , needs space to be indexed , hashed . So for a nice industry level uuid it is a must to be Memory Efficient , Should be lightning fast generated .
- First we shall talk about :
1. Memory Efficient : Best Industry standard UUIDs take 32(+4 hyphens) characters (Google UUID ) , 21 (a little new ) NanoID , But with slimUUID u gets it done in just **18 Characters** . Much better than industries old guys .
2. Time Efficiency : While creating any of these ID or tools much important thing to be considered is its benchmark testing . Its like a go to for comparing results , and dont be surprised we passed each and every test to come not just slight margins but at least **60% faster** in generation , hashing with comparison to industry best google UUID . For Exact numbers which are benchmark testing of UUIDV4 by google vs SlimUUID's various methods vs NanoID , and that are much better than the claims we did above , becuase we know talk is cheap [Refer to this Blog](https://dev.to/mitocondria/slimuuid-the-compact-memory-efficient-alternative-to-standard-uuids-2dak)

**TimeStamped** : Actually while storing the data in forms of rows , its beneficial to know when was this entry created , or else we need one more column to store this data .
- Solution : You need and we won't have done ... How is that possible . This is also provided in SlimUUID as you can retrace the time from UUID generated in form of date from the string slimUUID you have, this the magic of slimUUID's algorithm .

## Components of SlimUUID

SlimUUID  is made of total 18 characters (for Best and Fast Methods) : xxxxxxxxxx-xxxxxxxx

In these 18 chracters first 10 stands for timestamp till nanoseconds from a date that you can specify , but the time is not in integers rather , **Base-64 Encoding** . This Encoding is done using simple maths algorithms , which could be understood in time.go file .

Last 8 characters involves **xxHash** made from various parts which are listed : 
1. Mac Address / Can be IP address according to customization
2. Global Counter 
3. TimeSeed 
4. xxHash Sum64
5. TimeStamp from previous stage
Its arrangement can be seen in slimuuid.go file's different methods .

## Technical Details

**Character Set**:  
- **Base-64**: `0-9` (10 chars) + `a-z` (26 chars) + `A-Z` (26 chars) + `_,-`(2 chars) = **64 total**.

**Structure**:  
1. **Year** (1 char, 64 possible values, starting from 2025).  
2. **Month** (1 char, up to 12 possible values).  
3. **Day** (1 char, up to 31 possible values).  
4. **Hour** (1 char, up to 24 possible values).  
5. **Minute** (1 char, up to 60 possible values).  
6. **Millisecond** (2 chars, derived from current second + millisecond).  
7. **Nanosecond** (2 chars, derived from current millisecond + nanosecond).  
8. **xxHash** (8-12 chars depending upon methods being used).


## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/theMitocondria/slimuuid
   ```
# Project Setup and Usage

2. **Navigate into the project folder**:
   ```bash
   cd slimuuid
   ```

3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

4. **Build/Run** :
   ```bash
   go run .
   ```

## Usage
   ```
   func main() {
      /* function of your need and choice , can run go run doc to know each and every one of them */
       id := slimuuid.Generate()
       fmt.Println("Generated ID:", id)
   }
   ```

**Run the main program**:
   ```bash
   go run .
   ```


## Examples

### Basic ID Generation:
   ```go
   id := Generate()
   fmt.Println("My Unique ID:", id)
   // Output: e.g., 2mV5F7b6y0fHXJK9tz
   ```

## Contributing

Contributions, issues, and feature requests are welcome!

1. Fork the project repository.
2. Create a new branch:
   ```bash
   git checkout -b feature/YourFeature
   ```
3. Commit your changes:
   ```bash
   git commit -m 'Add some feature'
   ```
4. Push to the branch:
   ```bash
   git push origin feature/YourFeature
   ```
5. Open a Pull Request.

Do checkout blog for eact memory efficiency , hashing , benchmark testing reports .
## License
Distributed under the MIT License. See LICENSE for more information.


**Feel free to modify any sections to fit your project’s needs! If you have any questions or run into issues, open an issue or pull request on GitHub. Happy coding! 😊**
