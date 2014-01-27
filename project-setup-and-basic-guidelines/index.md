---
layout: default
title: Sass Basics -- Project Setup, Workflow, Nesting, etc
page-class: page--project-setup
---

<section class="section__block">
	<header class="section__header">
		<h2  class="section__title">Project Setup and Basic Guidelines</h2>
	</header>
	<div class="section__content">
		<h3  class="section__sub-title">Project Structure</h3>
		<p>The first thing to understand is that <code>.scss</code> is essentially a subset to the CSS3 syntac. This means that every valid CSS3 stylesheet is a valid <code>.scss</code> as well. These means that you can switch to <code>.scss</code> immediately once you have completed the installation. However, unlike vanilla CSS, Sass is a real scripting language with expressions, functions, variables, conditional logic, and loops. You don’t have to use all of these features to get some benefit out of Sass, but they’re there if you need them. Later we will get into how they will make complex and repetitive CSS much easier to write. </p>
		<p>Below is an example of the file structure we will be using in our projects. While at first this may seem very complicated, as you get more compfortable with <code>.scss</code> projects you will see that breaking the project up into simple files makes it much easier to maintain. </p>
		
{% highlight scss %}
styles/
|
|-- framework/             # Starting point for base styles
|   |-- _framework.scss    # Include to get access to all modules
|   |-- _defaults.scss     # Default variables for framework
|   |-- components/        # Optional files 
|       |-- _buttons.scss  # Etc...
|       ...
|   |-- core/              # The base of the framework
|       |-- _headings.scss # Etc...
|       ...
|   |-- generic/           # Tools to use in any project
|       |-- _mixins.scss   # Etc...
|       ...
|   ...
|
|-- cru/                   # Cru Styles and Base
|   |-- _cru.scss          # Include to get access to all modules
|   |-- _cru-colors.sass   # Etc...
|   |-- _cru-buttons.scss  # Etc... 
|   ...
|
|-- ui/                    # Project Specific UI
|   |-- _page.scss         # Etc...
|   |-- _page-header.scss  # Etc...
|   ...
|
`-- _vars.scss             # project specific variables
`-- main.scss              # primary Sass file
{% endhighlight %}	



<h3  class="section__sub-title">Partials</h3>	
<p>In the example above, you can see that the project is broken up into many different small files. While this is not good practice in CSS, due to the way Sass is compiled it is perfectly acceptable.  These files are called partials. Sass will compile each of these files into one large <code>.css</code> file based on the order the files are included. We will go into the ordering a little bit later.</p>
		
<p>Partials are to be included in their main parent folder. This means that the framework files will be imported into a <code>_framework.scss</code>, the cru files into a <code>_cru.scss</code> file, etc. We will then include these once more into the <code>_main.scss</code>. This allows for increased portability and ease of including the framework or cru files into any project.</p>
		
<p>The example below shows what our <code>main.scss</code> file will look like</p>
		
{% highlight scss %}
/**
 * Files to be included based on precident:
 *
 * Frameworks get included first, then project files
 */

/**
 * Project Variables
 */
@import "vars";

/**
 * Framework and Cru Styles
 */
@import "framework/framework";
@import "cru/cru";

/**
 * Project Files
 */
@import "ui/page";
@import "ui/page-header";
@import "ui/page-nav";
@import "ui/page-footer";

/* Overrides at the bottom */
@import "ui/utilities";
 ... 
{% endhighlight %}
		
<h3  class="section__sub-title">Source Order</h3>
		
<p>The example above demonstrated that the partials and projects should be imported in a specific order. A well ordered stylesheet will be ordered something like this:</p>
<ol>
	<li>Reset – ground zero.</li>
	<li>Elements – unclassed h1, unclassed ul etc.</li>
	<li>Objects and abstractions — generic, underlying design patterns.</li>
	<li>Components – full components constructed from objects and their extensions.</li>
	<li>Style trumps – error states etc.</li>
</ol>
			
<p>This means that—as you go down the document—each section builds upon and inherits sensibly from the previous one(s). There should be less undoing of styles, less specificity problems and all-round better architected stylesheets.</p>
		
<p>For further reading I cannot recommend Jonathan Snook’s <a href="http://smacss.com/">SMACSS</a> highly enough. below.</p>

</div>
</section>


<section class="section__block">
	<header class="section__header">
		<h2  class="section__title">Methodology</h2>
	</header>	
	<div class="section__content">
<h3  class="section__sub-title">General</h3>
<p>Limit your stylesheets to a maximum 80 character width where possible. Exceptions may be gradient syntax and URLs in comments. That’s fine, there’s nothing we can do about that.</p>
<p>I prefer four space size indentation and I write multi-line CSS.</p>
<h3  class="section__sub-title">Table of Contents</h3>
<p>At the top of stylesheets, a table of contents will detail the sections contained in the document, for example:</p>
{% highlight scss %}
/*------------------------------------*\
    $CONTENTS
\*------------------------------------*/
/**
 * CONTENTS............It's what you are reading
 * DEFAULTS............Project variables and settings
 * MIXINS..............Tools for our project
 * RESET...............Set our reset defaults
 */
{% endhighlight %}
<p>This will tell the next developer(s) exactly what they can expect to find in this file. Each item in the table of contents maps directly to a section title.</p>
<p>Since we will be working across multiple files then each item in the table of contents will map to an include which pulls that section in.</p>
		
<h3  class="section__sub-title">Section Titles</h3>
<p>The table of contents would be of no use unless it had corresponding section titles. Denote a section thus:</p>

{% highlight scss %}
/*------------------------------------*\
    $RESET
\*------------------------------------*/
{% endhighlight %}

<p>The <code>$</code> prefixing the name of the section allows us to run a find ([Cmd|Ctrl]+F) for $[SECTION-NAME] and limit our search scope to section titles only.</p>

<h3  class="section__sub-title">More on Partials</h3>
<p>Some people prefer to work with large, single files.This is fine, and by sticking to the following guidelines you’ll encounter no problems. For our purposes however, we will be breaking ours up into small files making it easier to maintain and work in collaboration.</p>

<h3  class="section__sub-title">Comments</h3>
<p>For the sake of uniformity, I propose a docBlock-eque commenting style which will be limited to 80 characters in length:</p>

{% highlight scss %}
/**
    * This is a docBlock style comment
    *
    * This is a longer description of the comment, describing the code in
    * detail. We limit these lines to a maximum of 80 characters in length.
    *
    * We can have markup in the comments, and should be encouraged to do so.
    *
    *
    */		
{% endhighlight %}
		
<h3  class="section__sub-title">Code Structure</h3>
<h4>Anatomy</h4>
<p>We will be following a number of standards when structuring rulesets.</p>
	<ul>
		<li>Use hyphen delimited class names (except for BEM notation, see below)</li>
		<li>4 space indented</li>
		<li>Multi-line</li>
		<li>Declarations in relevance (NOT alphabetical) order</li>
		<li>Indent vendor prefixed declarations so that their values are aligned</li>
		<li>Indent our rulesets to mirror the DOM</li>
	</ul>
<p>A brief example:</p>
			
{% highlight scss %}
.widget{
   /* Positioning */
   position: absolute;
   z-index: 10;
   top: 0;

   /* Display & Box Model */
   display: inline-block;
   overflow: hidden;
   width: 100px;
   border: 10px solid #333;
   margin: 10px;

   /* Color */
   background: #000;
   color: #fff;

   /* Text */
   font-family: sans-serif;
   font-size: 16px;

   /* Others */
   cursor: pointer;
   -webkit-border-radius: 4px;
      -moz-border-radius: 4px;
           border-radius: 4px;

}
    .widget-heading{
        margin-right: -10px;
        margin-left: -10px;
        padding: 0.25em;
		
        font-size: 1.5rem;
        line-height: 1;
        font-weight: bold;
    }
{% endhighlight %}
			
<p>We can see that .widget-heading must be a child of .widget as we have indented the .widget-heading ruleset one level deeper than .widget. This is useful information to developers that can now be gleaned just by a glance at the indentation of our rulesets.</p>
<p>One exception to the multi-line rule might be in cases of the following:</p>			

{% highlight scss %}
.t10    { width:10% }
.t20    { width:20% }
.t25    { width:25% }       /* 1/4 */
.t30    { width:30% }
{% endhighlight %}
			
<h4>Naming Conventions</h4>
<p>For the most part I recommend simply use hyphen delimited classes (e.g. .widget-heading, not .widget_heading or .widgetHeading), however in certain circumstances we should use BEM (Block, Element, Modifier) notation.</p>

<p>BEM is a methodology for naming and classifying CSS selectors in a way to make them a lot more strict, transparent and informative.</p>

<p>The naming convention follows this pattern:</p>
			
{% highlight scss %}
.block{}
.block__element{}
.block--modifier{}
{% endhighlight %}
<ul>
	<li><code>.block</code> represents the higher level of an abstraction or component.</li>
	<li><code>.block__element</code> represents a descendent of .block that helps form .block as a whole.</li>
	<li><code>.block--modifier</code> represents a different state or version of .block.</li>
</ul>
<p>An analogy of how BEM classes work might be:</p>

{% highlight scss %}
.person{}
.person--woman{}
    .person__hand{}
    .person__hand--left{}
    .person__hand--right{}
{% endhighlight %}
<p>BEM looks a little uglier, and is a lot more verbose, but it grants us a lot of power in that we can glean the functions and relationships of elements from their classes alone. Also, BEM syntax will typically compress (gzip) very well as compression favors/works well with repetition.</p>
<p>Regardless of whether we need to use BEM or not, always ensure classes are sensibly named; keep them as short as possible but as long as necessary. Ensure any objects or abstractions are very vaguely named (e.g. .ui-list, .media) to allow for greater reuse. Extensions of objects should be much more explicitly named (e.g. .user-avatar-link). Don’t worry about the amount or length of classes; gzip will compress well written code incredibly well.</p>	

<h4>OOCSS</h4>
<p>I work in an OOCSS manner; I split components into structure (objects) and skin (extensions). As an analogy (note, not example) take the following:</p>

{% highlight scss %}
.room{}

.room--kitchen{}
.room--bedroom{}
.room--bathroom{}
{% endhighlight %}

<p>We have several types of room in a house, but they all share similar traits; they all have floors, ceilings, walls and doors. We can share this information in an abstracted <code>.room{}</code> class. However we have specific types of room that are different from the others; a kitchen might have a tiled floor and a bedroom might have carpets, a bathroom might not have a window but a bedroom most likely will, each room likely has different colored walls. OOCSS teaches us to abstract the shared styles out into a base object and then extend this information with more specific classes to add the unique treatment(s).</p>
<p>So, instead of building dozens of unique components, try and spot repeated design patterns across them all and abstract them out into reusable classes; build these skeletons as base ‘objects’ and then peg classes onto these to extend their styling for more unique circumstances.</p>
<p>If you have to build a new component split it into structure and skin; build the structure of the component using very generic classes so that we can reuse that construct and then use more specific classes to skin it up and add design treatments.</p>


</div>	
</section>



