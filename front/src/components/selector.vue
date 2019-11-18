<template>
	<div class="selector">
		<form v-on:submit.prevent>
			<!-- Languages -->
			<country-flag :country="$language" />
			<select :value="$language" @change="$language = $event.target.value">
				<option>US</option>
				<option>FR</option>
				<option>DE</option>
				<option>IT</option>
			</select>

			<!-- Factions -->
			<select v-model="faction" @change="getRefs(faction, category)">
				<option v-for="f in factions" :key="f.id" :value="f.id">{{ f.name }}</option>
			</select>

			<!-- Categories -->
			<select v-model="category" @change="getRefs(faction, category)">
				<option v-for="c in categories" :key="c.id" :value="c.id">{{ c.name }}</option>
			</select>

			<!-- References -->
			<select v-model="ref">
				<option v-for="r in refs" :key="r.id" :value="r.id">[{{ r.status }}] #{{ r.id }} {{ r.title }}</option>
			</select>
			<button type="submit" @click="$emit('select_ref', ref)">Go</button>

			<!-- Create new reference -->
			<input v-model="newName" type="text" class="w-25" placeholder="new ref name" />
			<button type="submit" @click="createRef(newName)">Create</button>
		</form>
	</div>
</template>

<script>
import { Factions, Categories } from './const.js'
import { EventBus } from '../main.js'

export default {
	name: 'Selector',
	components: {},
	created: function() {
		EventBus.$on('refresh_selector', ref_id => {
			this.getRefs(this.faction, this.category)
			this.ref = ref_id
		})
		this.getRefs(this.faction, this.category)
	},
	data() {
		return {
			factions: Factions,
			faction: 11,
			categories: Categories,
			category: 5,
			refs: [],
			ref: 88,
			newName: '',
		}
	},
	methods: {
		getRefs: function(faction, category) {
			if (!faction || !category) {
				return
			}
			this.$http
				.get(
					process.env.VUE_APP_API_ENDPOINT +
						`/ref?faction_id=${faction}&category_id=${category}&lang=${this.$language}`
				)
				.then(function(res) {
					console.debug(res)
					this.refs = res.data
				})
				.catch(function(err) {
					console.error(err)
				})
		},
		createRef: function(name) {
			if (!this.faction || !this.category || !name) {
				return
			}
			var ref = {
				faction_id: this.faction,
				category_id: this.category,
				title: name,
			}
			this.$http
				.post(
					process.env.VUE_APP_API_ENDPOINT + `/ref?faction_id=${this.faction}&category_id=${this.category}`,
					ref
				)
				.then(function(res) {
					console.debug(res)
					this.$emit('select_ref', res.body)
					this.newName = ''
				})
				.catch(function(err) {
					console.error(err)
				})
		},
	},
}
</script>

<style lang="scss">
@import '../custom.scss';

.selector {
	@extend .row;
	form {
		@extend .form-inline;
		@extend .w-100;
	}
	select,
	input {
		@extend .form-control-sm;
		max-width: 30% !important;
	}
	button {
		@extend .btn-sm;
		@extend .mr-2;
	}
}
</style>
