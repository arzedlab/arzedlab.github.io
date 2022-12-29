+++
author = "Ravshan"
title = "Browser-in-the-Browser Attack"
date = "2022-12-29"
description = "Browser-in-the-browser(BitB) come to be known as a browser-in-the-browser attack"
tags = [
    "browser-in-the-browser",
    "bitb",
    "attack",
    "cybersecurity",
]
+++

Hello folks, My name is Ravshan. I am an enthusiast in the cybersecurity field. I like researching and learning in this field. I hope this topic will help you to be more aware.
The trend we know that cybercriminals are always inventing new ways to deceive users in order to obtain their credentials, secret keys, or other valuable information. In general, all these attacks are aimed at vigilant users, regardless of how sophisticated they become. The address of the website where you are being asked to enter your credentials is the first and foremost detail you must pay attention to so you can avoid phishing attacks. If you are not familiar with a particular domain, it is strongly recommended that you do not enter your credentials on that website. The next thing you need to pay attention to is the email address from which the message was sent and its format. Phishing emails usually have typos or grammatical errors in their text. These mistakes can help determine whether an email is real or fake. In addition, if there are hyperlinks in the message body, check them before clicking on them. 
In this article, we go through one the type of sophisticated phishing attacks as called Browser-in-the-Browser(BitB)

## What is a browser-in-the-browser attack?
Browser-in-the-browser(BitB) come to be known as a "browser-in-the-browser" attack was described by researcher mr.d0x. Nowadays, Creating websites is more powerful than we can imagine with HTML, CSS, Javascript, and other web development languages. It’s not hard to make a copy of any website and change the domain name and email address. So, if you receive an email from your bank asking for your login information, make sure it is sent from the official email address of your bank.
The concept behind a BitB attack is to build what seems to be a secure popup browser window created by the browser itself, but is actually nothing more than a web page displayed in an existing browser window. When we login to a website using Google, Microsoft, Apple, or another service, we are commonly served with a pop-up window asking us to authenticate. Here is an example of authentication to Canva via Apple Account.
![](/static/Clipboard_2022-12-29-16-21-12.png)

Unfortunately for us, even the vigilant user may be trapped by this schema. Because finding the difference between fake and real is quite difficult when viewing. But fortunately for cybercriminals, this technique is easy to build, it works like this: The first cybercriminal makes copies of some websites for example it can be Amazon.com, Etsy.com, Shopify any other website. Buys a domain name that is similar to the original if take Amazon as an example `amazonprime.xyz` or `amazoncheck.com` or `amazom.me`. Deploys phishing sites to hosting services and spam to confiding users. More and more users select to authenticate single-sign-on(SSO) via Google, Microsoft, or any other website. By reason of forgetting passwords and faster way of getting authenticated. Once users are authenticated, cybercriminals can access their data. For example, they can use your credit card information to make purchases on Amazon or buy things from other websites that you have an account with. Additionally, they can access all of the personal information related to your email accounts and social media profiles.
![](/static/Clipboard_2022-12-29-16-40-49.png)

## How can you tell if the login window is fake?

Real login windows look and behave like browser windows. You may maximize and minimize them, as well as move them around the screen. Fake pop-ups are tied to the page on which they appear. They may also move freely and cover buttons and graphics, but only inside their borders, i.e. the browser window. They are not permitted to go. That distinction should assist you in identifying them. Try moving the login window beyond the parent window's border. A genuine window will simply pass over; a fraudulent window will become trapped. And the last one, always use Two Factor Authentication. 

## Conclusion
The internet may be a frightening experience. Although cybercrime is a never-ending problem, you don't have to be scared of it if you follow all of the basic best practices, have your wits about you, and implement the proper security measures. Keeping up with scams and hacking tactics can at the very least keep you informed. Avoid being in a hurry, since taking your time reduces your chances of seeing what you assume is there rather than what is certainly there.
