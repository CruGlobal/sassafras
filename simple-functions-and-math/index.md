---
layout: default
title: Simple Functions and Math
page-class: page--simple-functions
---
<header class="page__header">
	<h1  class="page__title">{{page.title}}</h2>
</header>

<section  class="section__block">
	
<div class="section__content">
<h2 class="section__title">Built in Sass Functions</h2>

<p>Sass has a number of built in functions that will make our lives much easier. These can range from conversion of rgba colors to converting a unitless number to a percentage. We will only be covering the basics of functions and how to format them. For a full list of the built in functions visit the <a href="http://sass-lang.com/documentation/Sass/Script/Functions.html">Sass documentation</a>.</p>

<p>Using functions is a simple process as shown below. </p>

{% highlight scss %}
$base-ui-color: #cccccc;

a{
	color: darken($base-ui-color, 20%);
}
{% endhighlight %}

<p>The <code>.css</code> output</p>

{% highlight scss %}
a {
  color: #999999;
}
{% endhighlight %}

<p>This allows your code to become more dynamic than ever. For example if you chose to change the base color for your website all you need to do is change the variable and Sass will handle the conversions for you. </p>

<p>You also have the abillity to create your own functions which can, for example, convert pixel units to rem or em values. We will get into those a little later.</p>


</div>

<div class="section__content">
<h2 class="section__title">Operations and Math</h2>
<p>It was mentioned before that we can do some powerful things with variables. We were able to get a better view of this with functions but when we start implementing some operations we truely see how powerful they become.  We have the ability to plug in any numbers we wish but you will find that when used with variables, the abiliity to perform operations becomes very useful.  If set up correctly, changing one variable can result in an entirely different grid system or typography layout. </p>

<p>Here is a simple <code>.scss</code> file:</p>

{% highlight scss %}
$base-spacing-unit: 24px;

$gutter: $base-spacing-unit / 2;

.column{
  padding-left: $gutter;
  padding-left: $gutter / 16;
}
{% endhighlight %}

<p>And the <code>.css</code> output:</p>

{% highlight scss %}
.column {
  padding-left: 12px;
  padding-left: 0.75px;
}
{% endhighlight %}

<p>As you can see by altering one variable, we can changing the spacing of our girds and therefore, our entire horizontal layout.  As stated this was a very simple example of the types of math we can begin to do, as you will begin to see in future lessons. </p>

</div>

</section>