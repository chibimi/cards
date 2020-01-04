<template>
	<div>
		<template>
			<text-complete
				v-model="content"
				:rows="rows"
				:placeholder="placeholder"
				:strategies="strategies"
				@input="$emit('input', $event)"
			></text-complete>
		</template>
	</div>
</template>

<script>
import TextComplete from 'v-textcomplete'
import { ModelAdvantages,WeaponAdvantages,KnownWords } from './const.js'

const advantages = ModelAdvantages.concat(WeaponAdvantages).concat(KnownWords)
export default {
	name: 'TextArea',
	components: { TextComplete },
	props: ['value', 'rows', 'placeholder', 'abilities', 'ref_id'],
	watch: {
		ref_id: function() {
			this.strategies = this.getStategies()
		},
		value: function(newVal) {
			this.content = newVal
		},
	},
	created: function() {
		this.content = this.value
		this.strategies = this.getStategies()
	},
	data() {
		return {
			content: '',
			strategies: [],
		}
	},
	methods: {
		getStategies: function() {
			let _this = this
			return [
				{
					match: /(^|\s)#([a-zA-Z0-9+\-_]*)$/,
					template(name) {
						return '<span class="m-2">' + name.id + ' ' + name.title + '</span>'
					},
					search(item, callback) {
						callback(
							_this.abilities
								.filter(function(name) {
									if (name.title == null) {
										return false
									}
									return name.title.toLowerCase().startsWith(item.toLowerCase())
								})
								.slice(0, 10)
						)
					},
					replace(value) {
						return '$1#' + value.id + ':' + value.title + '#'
					},
				},
				{
					match: /(^|\s):([a-z0-9+\-_]*)$/,
					template(name) {
						return '<img width="17" src="advantages/' + name.label + '.jpg"></img> ' + name.label
					},
					search(item, callback) {
						callback(
							advantages
								.filter(function(name) {
									return name.label.toLowerCase().startsWith(item.toLowerCase())
								})
								.slice(0, 10)
						)
					},
					replace(value) {
						return '$1:' + value.label + ':'
					},
				},
			]
		},
	},
}
</script>

<style lang="scss" scoped>
@import '../custom.scss';
</style>
