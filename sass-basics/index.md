---
layout: default
title: Sass Basics
page-class: page--sass-basics
---
<header class="page__header">
	<h1  class="page__title">{{page.title}}</h2>
</header>
<section class="section__block">
<div class="section__content">
<h2  class="section__title">Nesting</h3>
<p>Nesting is a huge part of why you will love Sass, and most likely the first feature you will take advantage of when switching over from CSS. It allows you to write fewer lines of code in the end and helps you code to be much more readable. With Sass you can nest CSS rules inside each other instead of repeating selectors in separate declarations. The nesting will reflect the markup structure. </p>

<p>Let's look at a chunk of HTML: </p>

{% highlight html %}
 <header class="site-header" role="banner">
	 <h1>Let Them Eat Cake</h1>
 	 <ul  class="site-nav">
 	 	<li><a>Chocolate</a></li>
		<li><a>Oreo</a></li>
		<li><a>Carrot</a></li>
		...
 	 </ul>
	 ...
 </header>
{% endhighlight %}

<p>When writing Sass, I can mirror the element nesting, letting Sass handle building the full selector. Now lets look at the way we would write our <code>.scss</code> for the above piece of markup.</p>

{% highlight scss %}
.site-header{
	width: 100%;
	border-bottom: 1px solid #333333;
	
	margin-top: 16px;
	margin-top: 1rem;
	margin-bottom: 16px;
	margin-bottom: 1rem;
	
	h1{
		padding-top: 16px;
		padding-top: 1rem;
	}
}
{% endhighlight %}
	
<p>You will notice that I did not address all of the classes in our markup. This is because they do not need to be nested in the <code>.site-header</code>. In this case there will most likely only be one <code>.site-header</code> and <code>.site-nav</code>. Even though the <code>.site-nav</code> falls under the <code>.site-header</code> we do not need to nest them. However, there will be multiple <code>h1</code> elements throughout the page so more specificity is needed. </p>

<p>We would address both of these classes in this manner:</p>

{% highlight scss %}
.site-header{
	width: 100%;
	border-bottom: 1px solid #333333;
	
	margin-top: 16px;
	margin-top: 1rem;
	margin-bottom: 16px;
	margin-bottom: 1rem;
	
	h1{
		padding-top: 16px;
		padding-top: 1rem;
	}
}

    .site-nav{
	    width: 100%;
		
		padding-left: 24px;
		padding-left: 1.5rem;
		padding-right: 24px;
		padding-right: 1.5rem;
		
        /**
         * '>' means that only direct children of the will be affected. 
         */ 
		> li{
		    display: inline-block;
			
			> a{
			    color: #222222;		
			}
		}
	}
{% endhighlight %}

<p>By default Sass will organize the selectors in the manner that we want with a space between each selector. This is normally great, but there are a few times when this can make things difficult. For example, with psuedo classes. Lets use the <code>:hover</code> pseudo class as an an example.</p>

{% highlight scss %}
a{
	display: inline-block;
	color: #222222
	
	&:hover{
		text-decoration: none;
		color: #555555
	}
}
{% endhighlight %}

<p>By using the <code>&</code> symbol I am telling Sass to copy the above element and place it direclty in front of the <code>:hover</code></p>

<p>The <code>.css</code> output would be:</p>
{% highlight css %}
a{
	display: inline-block;
	color: #222222
}

a:hover{
	text-decoration: none;
	color: #555555
}
{% endhighlight %}

<p>This technique can be used frequently to reduce the size of you code but this is a little more advanced. </p>
</div>

<div class="section__content">
<h2 class="section__title">Variables</h2>
<p>Variables in CSS are a breakthrough in modern web development. With Sass you can create variables for things such as fonts, colors, margins, padding, and pretty much anything else you can think of. If you are familiar with PHP or JavaScript the concept is fairly transferable. </p>

<p>The best thing about variables is that they allow you to use an element more than ones, similar to a class in HTML. This will allow for much less copy and pasting and a more productive workflow.You can do many powerful things with variables, but for now we will deal with the basics.</p>

<p>For example, the following `.scss`:</p>

{% highlight scss %}
/**
 * Let's define some variables
 */
	$base-font-size:    16px;
	$base-line-height:  24px; 
	$base-ui-color:     #C0FFEE;

	$base-spacing-unit: $base-line-height;

/**
 * And apply them:
 */
.html{
	font-size: $base-font-size;
	line-height: $base-line-height;
	color: $base-ui-color;
}

.container{
	padding-left: $base-spacing-unit;
	padding-right: $base-spacing-unit;
}
{% endhighlight %}


<p>will output to this <code>.css</code>:</p>

{% highlight scss %}
.html {
    font-size: 16px;
    line-height: 24px;
    color: #c0ffee;
}

.container {
    padding-left: 24px;
    padding-right: 24px;
}
{% endhighlight %}

<p> This is mearly scrathing the surface of what variables can do. In the next lesson we will get into more advanced techniques and time saving applications.</p>
</div>

<div  class="section__content">
<h2  class="section__title">Methodology</h2>
<h3>Variable Placement</h3>
<p>Variable should, under most circumstances, be placed in a separate file just for variables and project settings. This makes it much easier to find a specific one in the future. This will typically be in a file with a naming converniton like: <code>_defaults.scss</code>, <code>_vars.scss</code>, or <code>_settings.scss</code>.  There may be a few reasons to place the variables within a certain file, but these are rare.</p>

<p>Under no circumstances should variables be placed within a css ruleset. This, by appearance, limits that variable to this ruleset. For the sake of keeping the code dry and readable place variable declarations at the top of your stylesheet and try to keep the naming as transferable as possible.</p>


<h3>More on Code Structure</h3>
<p>As talked about in the previous lesson we will be following a specific style when ordering our rules. We will continue adding on to these rules as we introduce new features. </p>


<h4>Nested Selectors Last</h4>
	{% highlight scss %}
	.weather {
	  background: #f7f7f7;

	  
	  > h3 {
	    border-bottom: 1px solid white;
	  }
	  
	}
	{% endhighlight %}
	
<p>Nothing goes after the nested stuff. And the same order previously stated will apply withing the nested selectors</p>
	
<h4>Maximum Nesting: Three Levels Deep</h4>

{% highlight scss %}
.weather {
    .cities{
		li{
			/* No More */
		}
	}
}
{% endhighlight %}

<p>I could go on and on about this but I will let <a href="http://css-tricks.com/">Chris Coyier</a> explain things:</p>
<blockquote>
	<p>Chances are, if you're deeper than that, you're writing a crappy selector. Crappy in that it's too reliant on HTML structure (fragile), overly specific (too powerful), and not very reusable (not useful). It's also on the edge of being difficult to understand.</p>
</blockquote>

<h4>Maximum Nesting: 50 Lines</h4>
<p>If a nested block of Sass is longer than that, there is a good chance it doesn't fit on one code editor screen, and starts becoming difficult to understand. The whole point of nesting is convenience and to assist in mental grouping. Don't use it if it hurts that.</p>



</div>
	

</section>