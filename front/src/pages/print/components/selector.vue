<template>
	<div class="ref">
		<div class="header">
			<h2 class="col-8">FR Card database</h2>
			<div class="col-4">
				<div class="float-right">
					<button v-on:click="reset()" class="form-control btn-sm btn-primary">Reset</button>
				</div>
			</div>
		</div>

		<div class="nav" role="tablist">
			<a v-for="f in factions" :key="f.id" data-toggle="tab" href="#nav" @click="faction = f">{{ f.name }}</a>
		</div>

		<div class="content">
			<div v-if="faction.id != null">
				<h3>{{ faction.name }}</h3>
				<div v-for="c in categories" :key="c.id">
					<Refs
						:card_ids="card_ids"
						:faction="faction.id"
						:category="c"
						v-on:add="add"
						v-on:remove="remove"
					/>
				</div>
			</div>
		</div>

		<div>
			<button v-on:click="generate_pdf" class="my-5 form-control btn-sm btn-primary">Generate PDF</button>
		</div>
	</div>
</template>

<script>
import { Factions, Categories } from '../../../const.js'
import Refs from './refs.vue'

export default {
	name: 'Selector',
	props: [],
	components: { Refs },
	watch: {},
	data() {
		return {
			card_ids: {},
			factions: Factions,
			categories: Categories,
			faction: {},
		}
	},
	methods: {
		reset: function () {
			this.card_ids = {}
		},
		add: function (n) {
			var oldval = this.card_ids[n]
			var newval = oldval != undefined ? oldval + 1 : 1
			this.$set(this.card_ids, n, newval)
		},
		remove: function (n) {
			var oldval = this.card_ids[n]
			if (!oldval) {
				return
			}
			var newval = oldval != undefined ? oldval - 1 : 0
			if (newval <= 0) {
				this.$delete(this.card_ids, n)
				return
			}
			this.$set(this.card_ids, n, newval)
		},
		generate_pdf: function () {
			console.log('passage')
			var cards = []
			for (const [key, value] of Object.entries(this.card_ids)) {
				for (var i = 0; i < value; i++) {
					cards.push(key)
				}
			}
			var url = process.env.VUE_APP_API_ENDPOINT + `/display?cards=${cards.join()}&lang=fr`

			var win = window.open(url, '_blank')
			win.focus()
			console.log(url)
		},
	},
}
</script>

<style lang="scss" scoped>
@import '../custom.scss';
.ref {
	.header {
		@extend .row;
		@extend .form-inline;
	}

	.error {
		@extend .row;
		@extend .alert;
		@extend .mx-0;
		@extend .py-1;
	}

	.nav {
		@extend .nav;
		@extend .nav-tabs;

		a {
			@extend .nav-item;
			@extend .nav-link;
		}
	}

	.content {
		// @extend .tab-content;
		@extend .container;
		@extend .px-0;
		@extend .mt-4;
	}
}
</style>
