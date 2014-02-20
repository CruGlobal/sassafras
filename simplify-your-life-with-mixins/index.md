---
layout: default
title: Simplify Your Life with Mixins
page-class: page--mixins
---
<header class="page__header">
	<h1  class="page__title">{{page.title}}</h2>
</header>

<section class="section__block">
<div class="section__content">
<h2 class="section__title">Mixins</h2>
<p>Of all of the Sass features, Mixins have to be the most helpful and most powerful. Mixins are akin to variables, if you pumped them full of steroids. You can define a complete style of an element, complete intensive calculations, and much more while having the ability to use them across your project.</p>

<p>Mixins are defined with the <code>@mixin</code> directive, which takes in a block of styles. You then apply this to a selectore by using the <code>@include</code> directive. </p>

<p>Lets start by creating a mixin that will center elements within ther parent. This does require that a <code>width</code> or <code>max-width</code> be set.</p>

{% highlight scss %}

@mixin center-block {
    display: block;
    margin-left: auto;
    margin-right: auto;
}

{% endhighlight %}

<p>Usage in <code>.scss</code></p>
{% highlight scss %}

.footer-wrap {
    width: 450px;
	
    @include center-block;
}

{% endhighlight %}

<p>Output in <code>.css</code></p>
{% highlight scss %}

.footer-wrap {
    width: 450px;
    display: block;
    margin-left: auto;
    margin-right: auto;
}

{% endhighlight %}


<p>Lets try something a little more complecated. This time we will be making a mixin that will Create vendor-prefixed CSS that are needed for some CSS3 declarations in one go.</p>


{% highlight scss %}
@mixin vendor($property, $value...){
    -webkit-#{$property}:$value;
       -moz-#{$property}:$value;
        -ms-#{$property}:$value;
         -o-#{$property}:$value;
            #{$property}:$value;
}
{% endhighlight %}

<p>Usage:</p>
{% highlight scss %}
.button{
	background-color: #cccccc;
	
	@include vendor(border-radius, 4px);
}
{% endhighlight %}

<p>Output:</p>

{% highlight scss %}
.button {
    background-color: #cccccc;
    -webkit-border-radius: 4px;
    -moz-border-radius: 4px;
    -ms-border-radius: 4px;
    -o-border-radius: 4px;
    border-radius: 4px;
}
{% endhighlight %}


<p>In this example, you will see that we wraped the <code>$property</code> variable with a <code>#{}</code>. This needs to be used whenever you want to use variables along with plain CSS in a mixin.</p>

<p>We will finish up with two of my favorite mixins:</p>


{% highlight scss %}
/**
 * Quickly generate a font-size in rems, with a pixel fallback, based on the
 * value we pass into the mixin:
 */
@mixin font-size($font-size, $line-height: true) {
    font-size: $font-size;
    font-size: ($font-size / $base-font-size) * 1rem;

    @if $line-height == true {
        line-height: ceil($font-size / $base-line-height) * ($base-line-height / $font-size);
    }

}

/**
 * Set a rem size and pixel fallback to any property of your choosing, based 
 * on the inputed values
 */
@mixin rem-calc($property, $size-value){
	#{$property}: $size-value;
	#{$property}: ($size-value / $base-font-size) * 1rem;
}
{% endhighlight %}


<p>Usage: </p>
{% highlight scss %}
h1{
	color: #333333;
	
	@include font-size(48px);
	@include rem-calc(padding-top, 24px);
}
{% endhighlight %}



<p>Output: </p>
{% highlight scss %}
h1 {
  color: #333333;
  font-size: 48px;
  font-size: 3rem;
  line-height: 1;
  padding-top: 24px;
  padding-top: 1.5rem;
}
{% endhighlight %}
</div>
<div class="section__content">
<h2 class="section__title">Methodology</h2>
<h3>More on Code Structure</h3>
<p>We are continuing in the process of defining our structure for writing our code.</p>
<p><code>@includes</code> should follow our regular styles but appear before nested selectors. This allows for greater readibility since they will all be grouped together at the bottom.</p>

<h4>Nested Selectors Last</h4>
	{% highlight scss %}
	.weather {
	  background: #f7f7f7;
	  
	  @include rem-calc(padding-left, 24px);
	  @include rem-calc(padding-right, 24px);
	  
	  > h3 {
	    border-bottom: 1px solid white;
		
		@include font-size(30px);
	  }
	  
	}
	{% endhighlight %}
	


</div>
</section>