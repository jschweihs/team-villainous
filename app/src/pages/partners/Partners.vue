<template>
    <div class="content max-600">
        <h1>Partners</h1>
        <ul class="partners">

            <li class="partner">
                <a href="https://sectorsixapparel.com/collections/villainous" target="_blank">
                    <img width="180" height="180" src="/images/partners/sectorsixapparel.jpg"/>
                </a>
                <p>We have partnered with Sector Six Apparel to design, manufacture, & distribute our extensive apparel store.  We are constantly designing & dropping new merchandise, so keep your eye on it!  Locate our store <a href="https://sectorsixapparel.com/collections/villainous" target="_blank">here</a>.</p>
            </li>

            <li class="partner">
                <a href="https://homg.net/" target="_blank">
                    <img width="180" height="180" src="/images/partners/hitormissgaming.jpg"/>
                </a>
                <p>We have partnered with Hit Or Miss Gaming to bring you top tier tournaments in Fortnite, Call of Duty, & other competitive titles.  Take a peek at HomG's website <a href="https://homg.net/" target="_blank">here</a>.</p>
            </li>

            <li class="partner">
                <a href="https://softgiving.com/" target="_blank">
                    <img width="180" height="180" src="/images/partners/softgiving.jpg"/>
                </a>
                <p>Our partnership with Softgiving provides our streamers with access to Charity Stream software which eases the burden of hosting a stream for charity.  Interested in working with Softgiving & hosting your own charity streams?  Find them <a href="https://softgiving.com/" target="_blank">here</a>.</p>
            </li>

            <li class="partner">
                <a href="https://advanced.gg/product-category/supplements/" target="_blank">
                    <img width="180" height="180" src="/images/partners/advanced_gg.jpg"/>
                </a>
                <p>AdvancedGG is revolutionizing the industry. They hold the first gaming supplement that uses premium ingredients that work. In addition, it's healthier and more cost-effective than most sugary energy drinks. Use code 'ViL' for 10% off <a href="https://advanced.gg/product-category/supplements/" target="_blank">here</a>.</p>
            </li>

        </ul>
    </div>
</template>

<style scoped>
    .partners {
        margin: 0 auto;
        padding: 0;
        list-style: none;
        max-width: 600px;
    }

    .partner {
        clear: both;
        margin-bottom: 10px;
        height: 180px;
        background-color: var(--color-grey-dark);
        border-top-left-radius: 20px;
        border-bottom-left-radius: 20px;

    }

    .partner img {
        float: left;
        max-height: 180px;
        margin-right: 20px;
        border-top-left-radius: 20px;
        border-bottom-left-radius: 20px;
    }

    .partner p {
        padding: 20px;
        height: 138px;
        margin-left: 180px;
        border-top: 1px solid var(--color-grey-light);
        border-right: 1px solid var(--color-grey-light);
        border-bottom: 1px solid var(--color-grey-light);
        border-top-right-radius: 20px;
        border-bottom-right-radius: 20px;
        background-color: var(--color-grey-medium);
    }

    .partner:last-child {
        margin-bottom: 0;
    }

    @media screen and (max-width: 700px) {
        .partner p {
            font-size: 12px;
        }
    }

    @media screen and (max-width: 450px) {
        .partner p {
            font-size: 10px;
        }
    }

    @media screen and (max-width: 380px) {
        .partner p {
            font-size: 9px;
        }
    }

</style>
