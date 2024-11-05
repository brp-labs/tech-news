<h3>The Latest News on AI and Machine Learning from Tech Xplore</h3>

<b>Author:</b> Brian Ravn Pedersen, Data Engineer and Software Developer<br/>
<b>Created:</b> 2024-11-05<br/>
<b>Technologies:</b> Go (Golang), React<br/>
<b>GitHub Repository:</b> https://github.com/brp-labs/tech-news<br/>

Data used for this application is sourced from <b>Tech Xplore</b> at https://techxplore.com/.<br>

This application makes a request to the RSS feed provided by Tech Xplore. This part is handled server-side by the Go file, which, after decoding the received XML-formatted RSS response into a Go structure, restructures the dataset into JSON format before outputting on port 8080.<br>

The React application listens on port 8080, grabs the data when it becomes available, formats the data within an HTML structure, and then outputs the final result to the DOM through port 3000.


<hr/>

<h3>Docker files will be included in this repository soon...</h3>