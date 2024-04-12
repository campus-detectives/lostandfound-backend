<a name="readme-top"></a>
<br />
<div align="center">
  <a href="https://github.com/github_username/repo_name">
    <img src="images/logo.png" alt="Logo" width="160" height="160">
  </a>

<h1 align="center">VITrace</h3>

  <p align="center">
    An app to help you trace your lost items
  </p>
</div>



<!-- ABOUT THE PROJECT -->
## About The Project
VITrace is a for the of lost and found items within campus environments. It offers a solution to address the challenges faced by both security personnel and students in locating and reclaiming lost belongings.

The process of handling lost items has traditionally relied on manual lists maintained by watchmen, resulting in inefficiencies and difficulties for owners trying to locate their lost possessions. With Vitrace, we introduce a centralized approach to streamline this process.

### Key Features
- Centralized System: VITrace replaces outdated manual lists by providing a centralized database accessible to authorized personnel.
- User-Friendly Interface: Our app offers an easy interface to upload images and obtains details of lost items.
- Object Detection: Our app offers an easy way to search for your lost items by uploading by searching for your app using an image.
- Ownership Tracking: VITrace records ownership information, ensuring transparency and aiding in conflict resolution when multiple claims are made for the same item.
- Section for Laundry Items: Recognizing the unique challenge of lost laundry items on campus, VITrace offers a section platform where students can independently upload and search for their missing items.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## How It Works

1. **Guard Uploads Found Item**:
   - When a guard finds a lost item, they use the VITrace app to upload an image of the item along with relevant details such as location and category.

2. **AI Model Generates Embedding for Image**:
   - Upon uploading, an AI model processes the image to generate an embedding. This embedding captures essential features and characteristics of the item.

3. **User Searches for Lost Item**:
   - Students or users can search for their lost items using the VITrace app based on various parameters such as location, category, and by uploading an image.

4. **Matching with Embeddings**:
   - VITrace compares the embeddings of the uploaded image (from the user's search) with the embeddings of found items stored in the database.
   - The system identifies potential matches based on similarity scores derived from the comparison of embeddings.

5. **Identifying Best Matches**:
   - Based on the comparison results, VITrace presents the user with a list of potential matches ranked by relevance.
   - Users can review the matched items and claim their belongings through the app.

6. **Conflict Resolution**:
   - VITrace tracks ownership information associated with claimed items to resolve any conflicts that may arise if multiple users claim the same item.


##

**Authority Side Workflow**

   <img src="./images/auth.jpeg">

**Student Side Workflow**

   <img src="./images/Student.jpeg">



## Getting Started
### Prerequisites

* golang-migrate
  ```sh
  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  ```
* make

### Configuration

The API app uses environment variables for configuration. You can set these variables directly in your environment or use a `.env` file placed at the project root.

#### Configuration Variables

The following configuration variables are used:

- **LAF_DSN**: PostgreSQL DSN (Data Source Name) to connect to the database server.

- **LAF_HTTP_ADDR**: The address on which the HTTP server will run. This should be specified in the form `host:port`.

### Setting Environment Variables

You can set these environment variables in your terminal or in a `.env` file. If using a `.env` file, make sure it is placed in the root directory of the project. An example `.env` file might look like this:

```
LAF_DSN=postgres://myusername:mypassword@localhost:5432/mydatabase
LAF_HTTP_ADDR=localhost:8000
```

### Usage
1. Clone the repo

2. Setup database
   ```sh
   make db/migrations/up
   ```

3. Run
   ```sh
   make run/api
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Screenshots

**Student Side**
<table>
  <tr>
    <td>
      <img src="images/screenshot_unclaimed.jpeg" alt="unclaimed items" height="400"/>
    </td>
    <td>
      <img src="images/screenshot_filter.jpeg" alt="filter" height="400"/>
    </td>
    <td>
      <img src="images/screenshot_result.jpeg" alt="filter" height="400"/>
    </td>
    <td>
      <img src="images/screenshot_contest_claim.jpeg" alt="filter" height="400"/>
    </td>
  </tr>
</table>

**Authority Side**
<table>
  <tr>
    <td>
      <img src="images/guard_home.jpeg" alt="unclaimed items" height="400"/>
    </td>
    <td>
      <img src="images/guard_submit.jpeg" alt="filter" height="400"/>
    </td>
    <td>
      <img src="images/guard_image.jpeg" alt="filter" height="400"/>
    </td>
    <td>
      <img src="images/guard_claim.jpeg" alt="filter" height="400"/>
    </td>
  </tr>
</table>


<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Team
- Heet Jhaveri - [Linkedin](https://www.linkedin.com/in/heet-jhaveri-873a90269/) | [GitHub](https://github.com/heet-10104) 
- Snehil Chatterjee - [Linkedin](https://www.linkedin.com/in/snehilchatterjee/) | [GitHub](https://github.com/snehilchatterjee)
- Amal Bangari - [Linkedin](https://www.linkedin.com/in/amal-bangari-692b29277/)  | [GitHub](https://github.com/amal064) 
- Vallari Ashar - [Linkedin](https://www.linkedin.com/in/vallari-ashar-500166249/) | [GitHub](https://github.com/heet-10104)
- Sarthak Bansal - [Linkedin](https://www.linkedin.com/in/sarthakbansal674/) | [GitHub](https://github.com/heet-10104)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
